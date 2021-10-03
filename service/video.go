package service

import (
	"BiliBili.com/cache"
	"BiliBili.com/model"
	"BiliBili.com/pkg/e"
	"BiliBili.com/pkg/utils"
	"BiliBili.com/serializer"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"mime/multipart"
	"strconv"
	"time"
)

type VideoShow struct {
	Title        string `json:"title" form:"title" bind:"required"`
	Video        string `json:"video"`
	VideoType    string `json:"video_type" form:"video_type"`
	Introduction string `json:"introduction" form:"introduction"`  	 //视频简介
	Uid          uint `json:"uid"`
	Author       model.User `json:"author"`
	Original     bool `json:"original"`        //是否为原创
	Weights      float32 `json:"weights"`     //视频权重
	Clicks       int `json:"clicks"`         //点击量
	Review       bool `json:"review"`   	 //是否审查通过
	PageSize int `json:"page_size" form:"page_size"`
	PageNum  int `json:"page_num" form:"page_num"`
}

type VideoRecommend struct {
	ID     uint   `json:"vid"`
	Title  string `json:"title"`
	Cover  string `json:"cover"`
	Author string `json:"author"`
	Clicks string `json:"clicks"`
}

type VideoInfo struct {
	Title  string `json:"title" form:"title" bind:"required"`
	Cover  string `json:"cover" form:"cover" bind:"required"`
	Introduction string `json:"introduction" form:"introduction" `
	Original bool `json:"original" form:"original" bind:"required"`
}

type VideoDelete struct {

}

type VideoInteractiveData struct {
	UID uint
}

type InteractiveData struct {
	Collect bool `json:"collect"`
	Like    bool `json:"like"`
	Follow  bool `json:"follow"`
}

func ClicksStoreInDB() {
	utils.Logfile("[info]", " Clicks are stored in the database")
	var vid int          //视频id
	var key string       //redis的key
	var clicks int       //点击量数字
	var strClicks string //字符串格式
	videos := cache.RedisClient.LRange(cache.ClicksVideoList, 0, -1).Val()
	for _, i := range videos {
		vid, _ = strconv.Atoi(i)
		key = cache.VideoClicksKey(vid)
		strClicks, _ = cache.RedisClient.Get(key).Result()
		clicks, _ = strconv.Atoi(strClicks)
		//删除redis数据
		cache.RedisClient.Del(key)
		//写入数据库
		model.DB.Model(&model.Video{}).Where("id = ?", vid).Update("clicks", clicks)
	}
	//删除list
	cache.RedisClient.Del(cache.ClicksVideoList)
	utils.Logfile("[info]", " Click volume storage completed")
}

// 收集点赞数和收藏数
func CollectAndLikeCount(vid string) (int, int) {
	var like int
	var collect int
	intVid,_ := strconv.Atoi(vid)
	strLike, _ := cache.RedisClient.Get(cache.VideoLikeKey(intVid)).Result()
	strCollect, _ := cache.RedisClient.Get(cache.VideoCollectKey(intVid)).Result()
	if strLike == "" || strCollect == "" {
		//like和SQL的关键词冲突了，需要写成`like`
		model.DB.Model(&model.Interactive{}).Where("vid = ? and `like` = 1", vid).Count(&like)
		model.DB.Model(&model.Interactive{}).Where("vid = ? and collect = 1", vid).Count(&collect)
		//写入redis，设置6小时过期
		cache.RedisClient.Set(cache.VideoLikeKey(intVid), like, time.Hour*6)
		cache.RedisClient.Set(cache.VideoCollectKey(intVid), collect, time.Hour*6)
		return like, collect
	}
	like, _ = strconv.Atoi(strLike)
	collect, _ = strconv.Atoi(strCollect)
	return like, collect
}

func GetClicksFromRedis(redis *redis.Client, vid int, dbClicks string) string {
	strClicks, _ := redis.Get(cache.VideoClicksKey(vid)).Result()
	if len(strClicks) == 0 {
		//将视频ID存入点击量列表
		redis.RPush(cache.ClicksVideoList, vid)
		//将点击量存入redis并设置25小时，防止数据当天过期
		redis.Set(cache.VideoClicksKey(vid), dbClicks, time.Hour*25)
		return dbClicks
	}
	return strClicks
}

func (service *VideoShow) Show (id string) serializer.Response{
	code := e.SUCCESS
	var video model.Video
	model.DB.Model(&model.Video{}).Preload("Author").
		Where("id = ? And review = true",id).First(&video)
	if video.ID == 0 {
		code = e.InvalidParams
		return serializer.Response{
			Status:code,
			Msg:e.GetMsg(code),
			Data:"视频不见了！",
		}
	}
	vid,_ := strconv.Atoi(id)
	like,count := CollectAndLikeCount(id)
	strClicks, _ := cache.RedisClient.Get(cache.VideoClicksKey(vid)).Result()
	if strClicks == "" {
		cache.RedisClient.RPush(cache.ClicksVideoList, vid)
		cache.RedisClient.Set(cache.VideoClicksKey(vid), video.Clicks, time.Hour*25)
	}
	cache.RedisClient.Incr(cache.VideoClicksKey(vid))
	data := serializer.VideoData{
		LikeCount:like,
		CollectCount:count,
	}
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data: serializer.BuildVideo(video,data),
	}
}

func (service *VideoRecommend) Recommend() serializer.Response {
	code := e.SUCCESS
	var videos []model.Video
	var count int
	model.DB.Model(&model.Video{}).Where("review = 1").Order("click desc").
		Find(&videos).Count(&count)
	for i:=0;i<count;i++ {
		tmp := strconv.Itoa(videos[i].Clicks)
		click ,_ := strconv.Atoi(GetClicksFromRedis(cache.RedisClient, int(videos[i].ID),tmp))
		videos[i].Clicks = click
	}
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data:serializer.BuildVideos(videos),
	}
}

func (service *VideoShow) List(id string) serializer.Response {
	code := e.SUCCESS
	var videos []model.Video
	var count int
	if service.PageSize == 0 {
		service.PageSize=10
	}
	model.DB.Model(model.Video{}).Where("uid = ?",id).Count(&count).
		Limit(service.PageSize).Offset((service.PageNum-1)*service.PageSize).
		Find(&videos)

	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data:serializer.BuildListResponse(serializer.BuildVideos(videos),uint(count)),
	}
}

func (service *VideoShow) Favor(id string) serializer.Response {
	code := e.SUCCESS
	var favorite []model.Interactive
	var count int
	model.DB.Model(model.Interactive{}).Where("uid = ? AND collect = true",id).
	Count(&count).Limit(service.PageSize).Offset((service.PageSize-1)*service.PageNum).
		Preload("Video").Find(&favorite)
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data:serializer.BuildListResponse(serializer.BuildFavors(favorite),uint(count)),
	}
}

func (service *VideoInfo) Update(id string,uid uint) serializer.Response {
	code := e.SUCCESS
	var video model.Video
	model.DB.Where(model.Video{}).Where("id = ?",id).First(&video)
	video.Title = service.Title
	video.Introduction = service.Introduction
	video.Original = service.Original
	model.DB.Save(&video)
	err := model.DB.Model(&model.Video{}).Where("id = ? and uid = ?", id, uid).
		Updates(map[string]interface{}{"title": service.Title,
			"introduction": service.Introduction, "original": service.Introduction}).Error
	if err!=nil {
		code = e.ERROR
		return serializer.Response{
			Status:code,
			Msg:e.GetMsg(code),
			Data:"视频更新数据出错",
		}
	}
	err = model.DB.Model(&model.Review{}).Where("vid = ?", id).
		Updates(map[string]interface{}{"status": 1000}).Error
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status:code,
			Msg:e.GetMsg(code),
			Data:"修改审核状态失败",
		}
	}
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data:"更新信息成功，重新进入审核",
	}
}

func (service *VideoDelete) Delete(id string) serializer.Response {
	code := e.SUCCESS
	var video model.Video
	model.DB.Where(model.Video{}).Where("id=?",id).Delete(&video)
	idT, _ := strconv.Atoi(id)
	cache.RedisClient.Del(cache.VideoClicksKey(idT))
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data:"删除成功",
	}
}

func (service *VideoShow) Upload(id uint,file,cover multipart.File,coverSize,fileSize int64) serializer.Response {
	code := e.SUCCESS
	status , info := utils.UploadToQiNiu(file,fileSize)
	_ , coverUrl := utils.UploadToQiNiu(cover,coverSize)
	if status != 200 {
		return serializer.Response{
			Status:  status  ,
			Data:      e.GetMsg(status),
			Error:info,
		}
	}
	fmt.Println("id",id)
	video := model.Video {
		Title : service.Title,
		Cover : coverUrl,
		Introduction : service.Introduction,
		Original : service.Original,
		Uid : id,
		Video:info,
		VideoType:viper.GetString("server.coding"),
	}
	model.DB.Model(model.Video{}).Create(&video)  // 创建视频
	model.DB.Create(&model.Review{Vid: video.ID, Status: 500}) // 创建审核视频
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data:"上传成功",
	}
}

func (service *VideoInteractiveData)Show(vid string,cid uint) serializer.Response{
	code := e.SUCCESS
	var interactive model.Interactive
	var follow model.Follow
	var count int
	var fans bool
	model.DB.Model(&model.Interactive{}).Where("vid=? AND uid = ?",vid,cid).First(&interactive) // 找到互动信息
	model.DB.Model(&model.Follow{}).Where("uid = ? AND cid = ?",interactive.Video.Uid,cid).First(&follow).Count(&count)
	if count == 1 {
		fans =true
	}
	data := InteractiveData {
		Collect:interactive.Collect,
		Like:interactive.Like,
		Follow:fans,
	}
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data: data,
	}

}
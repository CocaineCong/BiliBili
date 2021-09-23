package service

import (
	"BiliBili.com/cache"
	"BiliBili.com/model"
	"BiliBili.com/pkg/e"
	"BiliBili.com/pkg/utils"
	"BiliBili.com/serializer"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

type VideoShow struct {
	Title        string
	Cover        string
	Video        string
	VideoType    string
	Introduction string  	 //视频简介
	Uid          uint
	Author       model.User
	Original     bool        //是否为原创
	Weights      float32     //视频权重
	Clicks       int         //点击量
	Review       bool   	 //是否审查通过
}

type VideoRecommend struct {
	ID     uint   `json:"vid"`
	Title  string `json:"title"`
	Cover  string `json:"cover"`
	Author string `json:"author"`
	Clicks string `json:"clicks"`
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

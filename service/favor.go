package service

import (
	"BiliBili.com/cache"
	"BiliBili.com/model"
	"BiliBili.com/pkg/e"
	"BiliBili.com/serializer"
	"strconv"
)

type FavorCreateService struct {

}

type FavorDeleteService struct {

}

type LikeCreateService struct {

}

type LikeDeleteService struct {

}

func (service *FavorCreateService) Create(vid string,uid uint) serializer.Response {
	code := e.SUCCESS
	var data model.Interactive
	model.DB.Model(&model.Interactive{}).Where("uid = ? AND vid = ?",uid,vid).First(&data)
	if data.Collect == true {  // 已经收藏了
		code = e.ErrorFavorExist
		return serializer.Response{
			Status:code,
			Msg:e.GetMsg(code),
		}
	}
	intVid,_ := strconv.Atoi(vid)
	model.DB.Model(&model.Interactive{}).Where("uid = ? AND vid = ?", uid, vid).Update("collect", true)
	strCollect, _ := cache.RedisClient.Get(cache.VideoCollectKey(intVid)).Result()
	if strCollect != "" {
		cache.RedisClient.Incr(cache.VideoCollectKey(intVid))
	}
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data:"收藏成功",
	}
}

func (service *FavorDeleteService) Delete(vid string,uid uint) serializer.Response {
	code := e.SUCCESS
	var favor model.Interactive
	model.DB.Model(model.Interactive{}).Where("vid = ? AND uid = ?",vid,uid).First(&favor)
	if favor.Collect!=true {
		code = e.ErrorFavorExist
		return serializer.Response{
			Status:code,
			Msg:e.GetMsg(code),
			Data:"收藏不存在",
		}
	}
	model.DB.Model(&model.Interactive{}).Where("uid = ? AND vid = ?", uid, vid).Update("collect", false)
	intVid,_ := strconv.Atoi(vid)
	strCollect, _ := cache.RedisClient.Get(cache.VideoCollectKey(intVid)).Result()
	if strCollect != "" {
		cache.RedisClient.Decr(cache.VideoCollectKey(intVid))
	}
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),

	}
}

func (service *LikeCreateService) Create(vid string,uid uint) serializer.Response {
	code := e.SUCCESS
	var data model.Interactive
	model.DB.Model(&model.Interactive{}).Where("vid = ? AND uid = ?",vid,uid).First(&data)
	if data.Like == true {
		code = e.ErrorLikeExist
		return serializer.Response{
			Status:code,
			Msg:e.GetMsg(code),
			Data:"已经赞过啦！",
		}
	}
	model.DB.Model(&model.Interactive{}).Where("uid = ? AND vid = ?", uid, vid).Update("like", true)
	intVid,_ := strconv.Atoi(vid)
	strLike, _ := cache.RedisClient.Get(cache.VideoLikeKey(intVid)).Result()
	if strLike != "" {
		cache.RedisClient.Incr(cache.VideoLikeKey(intVid))
	}
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
	}
}

func (service *LikeDeleteService) Delete(vid string,uid uint) serializer.Response{
	code := e.SUCCESS
	model.DB.Model(&model.Interactive{}).Where("uid = ? AND vid = ?", uid, vid).Update("like", false)
	intVid,_ := strconv.Atoi(vid)
	strLike, _ := cache.RedisClient.Get(cache.VideoLikeKey(intVid)).Result()
	if strLike != "" {
		cache.RedisClient.Incr(cache.VideoLikeKey(intVid))
	}
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
	}
}
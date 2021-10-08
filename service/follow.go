package service

import (
	"BiliBili.com/model"
	"BiliBili.com/pkg/e"
	"BiliBili.com/serializer"
)

type ListFollowingService struct {

}

type ListFollowerService struct {

}


func (service *ListFollowingService)List(id string)serializer.Response {
	code := e.SUCCESS
	var followList []model.Follow
	var followingInfo []model.User
	model.DB.Model(&model.Follow{}).Where("fid = ?",id).Find(&followList)
	for _,follower := range followList{
		var following model.User
		model.DB.Model(&model.User{}).Where("id = ?",
			follower.ID).First(&following)
		followingInfo = append(followingInfo, following)
	}
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data: serializer.BuildListResponse(serializer.BuildUsers(followingInfo),uint(len(followingInfo))),
	}
}

func (service *ListFollowerService)List(id string)serializer.Response {
	code := e.SUCCESS
	var followList []model.Follow
	var followingInfo []model.User
	model.DB.Model(&model.Follow{}).Where("uid = ?",id).Find(&followList)
	for _,follower := range followList{
		var following model.User
		model.DB.Model(&model.User{}).Where("id = ?",
			follower.ID).First(&following)
		followingInfo = append(followingInfo, following)
	}
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data: serializer.BuildListResponse(serializer.BuildUsers(followingInfo),uint(len(followingInfo))),
	}
}
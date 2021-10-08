package service

import (
	"BiliBili.com/model"
	"BiliBili.com/pkg/e"
	"BiliBili.com/serializer"
	"strconv"
)

type ListFollowingService struct {

}

type ListFollowerService struct {

}

type CreateFollowingService struct {

}

type DeleteFollowingService struct {

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

func (service *CreateFollowingService)Create(id string,uid uint)serializer.Response {
	code := e.SUCCESS
	idInt ,_ := strconv.Atoi(id)
	follow:=model.Follow{
		Fid:uid,
		Uid:uint(idInt),
	}
	model.DB.Create(&follow)
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data:"关注成功",
	}
}

func (service *DeleteFollowingService)Delete(id string,uid uint)serializer.Response {
	code := e.SUCCESS
	var follower model.Follow
	idInt ,_ := strconv.Atoi(id)
	model.DB.Model(&model.Follow{}).Where("fid = ? AND uid = ?",
		uid,idInt).Delete(&follower)
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data:"取关成功",
	}
}
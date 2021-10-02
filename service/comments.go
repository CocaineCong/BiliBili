package service

import (
	"BiliBili.com/model"
	"BiliBili.com/pkg/e"
	"BiliBili.com/serializer"
	"strconv"
)

type ShowCommentService struct {
	PageSize int `json:"page_size" form:"page_size"`
	PageNum int `json:"page_num" form:"page_num"`
}

type DeleteCommentService struct {

}

type DeleteReplyService struct {

}

type CreateCommentService struct {
	Content string `json:"content" form:"content" bind:"required"`
}

type CreateReplyService struct {
	Uid uint `json:"uid" form:"uid" bind:"required"`
	Content string `json:"content" form:"content" bind:"required"`
}


func (service *ShowCommentService) Show(id string) serializer.Response {
	code := e.SUCCESS
	var count uint
	var comments []model.CommentStrut
	model.DB.Model(&model.Comment{}).Where("vid = ?",id).Count(&count)
	model.DB.Model(&model.Comment{}).Limit(service.PageSize).Offset((service.PageNum-1)*(service.PageSize)).
		Find(&comments)
	for i:=0;i<len(comments);i++{
		model.DB.Model(&model.Reply{}).Where("cid=?",comments[i].ID).Find(&comments[i].Reply)
	}
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data:serializer.BuildListResponse(serializer.BuildComments(comments),count),
	}
}

func (service *DeleteCommentService) Delete(cid string,uid uint) serializer.Response {
	code := e.SUCCESS
	model.DB.Where("id = ? and uid = ?",cid,uid).Delete(model.Comment{})
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data:"删除评论成功",
	}
}

func (service *DeleteReplyService) Delete(id string, uid uint) serializer.Response {
	code := e.SUCCESS
	model.DB.Where("id = ? and uid = ?",id,uid).Delete(model.Reply{})
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data:"删除回复成功",
	}
}

func (service *CreateCommentService) Create(vid string,uid uint)serializer.Response {
	code := e.SUCCESS
	vidInt ,_ := strconv.Atoi(vid)
	comment:=model.Comment{
		Vid:uint(vidInt),
		Content:service.Content,
		Uid:uid,
	}
	model.DB.Create(&comment)
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data:"创建评论成功",
	}
}

func (service *CreateReplyService) Create(cid string,uid uint)serializer.Response {
	code := e.SUCCESS
	var user model.User
	cidInt ,_ := strconv.Atoi(cid)
	model.DB.First(&user,uid)
	reply:=model.Reply{
		Cid:       uint(cidInt),
		ReplyUid:  uid,
		ReplyName: user.UserName,
		Content:   service.Content,
		Uid: service.Uid      ,
	}
	model.DB.Create(&reply)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   "创建回复成功",
	}
}
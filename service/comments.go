package service

import (
	"BiliBili.com/model"
	"BiliBili.com/pkg/e"
	"BiliBili.com/serializer"
)

type ShowCommentService struct {
	PageSize int `json:"page_size" form:"page_size"`
	PageNum int `json:"page_num" form:"page_num"`
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

package service

import (
	"BiliBili.com/model"
	"BiliBili.com/pkg/e"
	"BiliBili.com/serializer"
	"strconv"
)

type ListDanmuService struct {

}

type CreateDamuService struct {
	Vid uint `json:"vid" form:"vid"`
	Time  uint   `json:"time" form:"time"`
	Type  int    `json:"type" form:"type"`
	Color string `json:"color" form:"color"`
	Text  string `json:"text" form:"text"`
}

func (service *ListDanmuService)List(vid string) serializer.Response{
	code := e.SUCCESS
	var danmuList []model.Danmu
	model.DB.Model(model.Danmu{}).Where("vid = ?",vid).Find(&danmuList)
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data:serializer.BuildDanmus(danmuList),
	}
}

func (service *CreateDamuService)Create(uid string) serializer.Response {
	code := e.SUCCESS
	id,_ := strconv.Atoi(uid)
	danmu := model.Danmu{
		Vid:   service.Vid,
		Time:  service.Time,
		Type:  service.Type,
		Color: service.Color,
		Text:  service.Text,
		Uid:   uint(id),
	}
	model.DB.Create(&danmu)
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data:"创建成功啦！",
	}
}


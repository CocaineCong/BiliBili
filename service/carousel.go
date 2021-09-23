package service

import (
	"BiliBili.com/model"
	"BiliBili.com/pkg/e"
	"BiliBili.com/serializer"
)

type Carousel struct {
	Img string `json:"img"`
	Url string `json:"url"`
}

func (service *Carousel) Carousel() serializer.Response{
	code := e.SUCCESS
	var carousels []model.Carousel
	model.DB.Model(&model.Carousel{}).Find(&carousels)
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data:serializer.BuildCarousels(carousels),
	}
}

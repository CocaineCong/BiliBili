package serializer

import "BiliBili.com/model"

type Carousel struct {
	Id uint `json:"id"`
	Img   string   `json:"img"`
	Url   string `json:"img_path"`
}

func BuildCarousel(item model.Carousel) Carousel {
	return Carousel{
		Id:    item.ID,
		Img:   item.Img,
		Url:   item.Url,
	}
}

func BuildCarousels(items []model.Carousel) (carousels []Carousel) {
	for _, item := range items {
		carousel := BuildCarousel(item)
		carousels = append(carousels, carousel)
	}
	return carousels
}
package serializer

import "BiliBili.com/model"

type VideoFavor struct {
	ID    uint   `json:"vid"`
	Title string `json:"title"`
	Cover string `json:"cover"`
}

func BuildFavor(item model.Interactive) VideoFavor {
	return VideoFavor{
		ID:item.Vid,
		Title:item.Video.Title,
		Cover:item.Video.Cover,
	}
}

func BuildFavors(items []model.Interactive) (favors []VideoFavor) {
	for _,item := range items {
		favor := BuildFavor(item)
		favors=append(favors, favor)
	}
	return favors
}


package serializer

import (
	"BiliBili.com/cache"
	"BiliBili.com/model"
	"time"
)

type Video struct {
	Id uint `json:"id"`
	Title string `json:"title"`
	Cover string `json:"cover"`
	Video string `json:"video"`
	VideoType string `json:"video_type"`
	Introduction string `json:"introduction"`
	CreateAt time.Time `json:"create_at"`
	Original bool `json:"original"`
	Author User `json:"author"`
	Data VideoData `json:"data"`
	Clicks string `json:"clicks"`
}

type VideoData struct {
	LikeCount    int `json:"like_count"`
	CollectCount int `json:"collect_count"`
}

func BuildVideo(item model.Video,data VideoData) Video {
	clicks, _ := cache.RedisClient.Get(cache.VideoClicksKey(int(item.ID))).Result()
	return Video{
		Id:item.ID,
		Title:item.Title,
		Cover:item.Cover,
		Video:item.Video,
		VideoType:item.VideoType,
		Introduction:item.Introduction,
		CreateAt:item.CreatedAt,
		Original:item.Original,
		Author: User{
			ID:     item.Author.ID,
			UserName: item.Author.UserName,
			Sign:   item.Author.Sign,
			Avatar: item.Author.Avatar,
		},
		Data:data,
		Clicks:clicks,
	}
}

func BuildVideos(items []model.Video) (videos []Video) {
	data := VideoData{}
	for _,item := range items{
		video := BuildVideo(item,data)
		videos=append(videos, video)
	}
	return videos
}
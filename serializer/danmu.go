package serializer

import "BiliBili.com/model"

type Damu struct {
	Time  uint   `json:"time" form:"time"`
	Type  int    `json:"type" form:"type"`
	Color string `json:"color" form:"color"`
	Text  string `json:"text" form:"text"`
}

func BuildDanmu(item model.Danmu) Damu {
	return Damu{
		Time:item.Time,
		Type:item.Type,
		Color:item.Color,
		Text:item.Text,
	}
}

func BuildDanmus(items []model.Danmu) (danmus []Damu) {
	for _,item := range items {
		danmu := BuildDanmu(item)
		danmus = append(danmus, danmu)
	}
	return danmus
}



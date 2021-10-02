package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	Vid     uint   `gorm:"not null;index"`             //视频ID
	Content string `gorm:"type:varchar(255);not null"` //内容
	Uid     uint   `gorm:"not null"`                   //用户
}


type CommentStrut struct {
	ID uint `json:"cid"`//评论ID
	CreatedAt time.Time `json:"created_at"`
	Content string `json:"content"`//内容
	Uid uint `json:"uid"`
	Name string `json:"name"`
	Avatar string `json:"avatar"`
	Reply []ReplyStrut `json:"reply"`
}

type ReplyStrut struct {
	ID uint `json:"rid"`//回复id
	CreatedAt time.Time `json:"created_at"`
	Content string `json:"content"`//内容
	Uid uint `json:"uid"`
	Name string `json:"name"`
	Avatar string `json:"avatar"`
	ReplyUid  uint   `json:"reply_uid"`  //回复的人的uid
	ReplyName string `json:"reply_name"` //回复的人的昵称
}

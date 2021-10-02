package model

import "github.com/jinzhu/gorm"

type Reply struct {
	gorm.Model
	Cid       uint   `gorm:"not null;index"`             //评论的ID
	Content   string `gorm:"type:varchar(255);not null"` //内容
	Uid       uint   `gorm:"not null"`                   //评论的用户ID
	ReplyUid  uint   //回复的人的uid
	ReplyName string `gorm:"type:varchar(20);"` //回复的人的昵称
}
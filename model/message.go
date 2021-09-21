package model

import "github.com/jinzhu/gorm"

type Messages struct {
	gorm.Model
	Uid uint `gorm:"not null;"` //用户ID
	Fid uint  `gorm:"not null;"`//关联ID
	FromId uint `gorm:"not null;"`// 发送者
	ToId uint `gorm:"not null;"` // 接受者
	Content string `gorm:"size:255;"`
	Status int `gorm:"default:0"`//已读状态
}

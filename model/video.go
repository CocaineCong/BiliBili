package model

import "github.com/jinzhu/gorm"

type Video struct {
	gorm.Model
	Title        string  `gorm:"type:varchar(50);not null;index"`
	Cover        string  `gorm:"size:255;not null"`
	Video        string  `gorm:"size:255"`
	VideoType    string  `gorm:"varchar(5)"`
	Introduction string  `gorm:"varchar(100);default:'什么也没有'"` //视频简介
	Uid          uint    `gorm:"not null;index"`
	Author       User    `gorm:"ForeignKey:id;AssociationForeignKey:uid"`
	Original     bool    `gorm:"not null"`      //是否为原创
	Weights      float32 `gorm:"default:0"`     //视频权重
	Clicks       int     `gorm:"default:0"`     //点击量
	Review       bool    `gorm:"default:false"` //是否审查通过
}

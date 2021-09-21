package model

import "github.com/jinzhu/gorm"

type Announce struct {
	gorm.Model
	Title   string `gorm:"type:varchar(50);not null"`
	Content string `gorm:"type:varchar(200);"`
	Url     string `gorm:"type:varchar(100);"`
}

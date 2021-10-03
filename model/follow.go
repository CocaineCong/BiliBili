package model

import "github.com/jinzhu/gorm"

type Follow struct {
	gorm.Model
	Uid uint `gorm:"not null"`
	Fid uint `gorm:"not null"`
}

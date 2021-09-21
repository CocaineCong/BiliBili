package model

import "github.com/jinzhu/gorm"

type Admin struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);"`
	Email    string `gorm:"varchar(20);not null"`
	Password string `gorm:"size:255;not null"`
	Authority int `gorm:"not null"`
}

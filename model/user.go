package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Avatar   string    `gorm:"size:255;"`
	UserName     string    `gorm:"type:varchar(20);not null"`
	Email    string    `gorm:"varchar(20);not null;index"`
	Password string    `gorm:"size:255;not null"`
	Gender   int       `gorm:"default:0"`
	Authority int  `gorm:"not null"`
	Birthday time.Time `gorm:"default:'2019-09-05'"`
	Sign     string    `gorm:"varchar(50);default:'这个人很懒，什么都没有留下'"`
	Relations  []User `gorm:"many2many:relation; association_jointable_foreignkey:relation_id"`
}
package model

import "github.com/jinzhu/gorm"

type Review struct {
	gorm.Model
	Vid     uint   `gorm:"not null;index"` //视频ID
	Video   Video  `gorm:"ForeignKey:id;AssociationForeignKey:vid"`
	Status  int    `gorm:"not null"`          //审核状态
	Remarks string `gorm:"type:varchar(20);"` //备注
}

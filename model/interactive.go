package model

import "github.com/jinzhu/gorm"

type Interactive struct {
	gorm.Model
	Uid     uint `gorm:"not null"`
	Vid     uint `gorm:"not null"`
	Collect bool `gorm:"default:false"` //是否收藏
	//like和SQL的关键词冲突了，查询时需要写成`like`
	Like  bool  `gorm:"default:false"` //是否点赞
	Video Video `gorm:"ForeignKey:id;AssociationForeignKey:vid"`
}

package model

import "github.com/jinzhu/gorm"

type GoodsCarousel struct {
	gorm.Model
	BaseModel
	GoodsInfoId uint `gorm:"int(10);index"`
	Src         uint `gorm:"size:128"`
}

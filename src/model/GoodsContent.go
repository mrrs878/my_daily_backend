package model

import "github.com/jinzhu/gorm"

type GoodsContent struct {
	gorm.Model
	BaseModel
	GoodsInfoId uint   `gorm:"int(10);index"`
	Src         string `gorm:"varchar(128)"`
}

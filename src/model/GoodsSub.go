package model

import "github.com/jinzhu/gorm"

type GoodsSub struct {
	gorm.Model
	BaseModel
	GoodsInfoId  uint   `gorm:"int(10);index"`
	Label        string `gorm:"varchar(64)"`
	Value        string `gorm:"varchar(128)"`
	Price        string `gorm:"varchar(16)"`
	PromotePrice string `gorm:"varchar(16)"`
	Weight       string `gorm:"varchar(8)"`
	Volume       uint   `gorm:"int(16)"`
	Stock        uint   `gorm:"int(16)"`
}

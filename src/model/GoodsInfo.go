package model

import "github.com/jinzhu/gorm"

type Goods struct {
	gorm.Model
	BaseModel
	Name        string `gorm:"varchar(64)"`
	Service     string `gorm:"varchar(64)"`
	Class       uint   `gorm:"int(10);index"`
	Description string `gorm:"varchar(64)"`
	TotalVolume uint   `gorm:"int(10)"`
	TotalSee    uint   `gorm:"int(10)"`
	TotalStock  uint   `gorm:"int(10)"`
}

package model

import "github.com/jinzhu/gorm"

type DataDict struct {
	gorm.Model
	BaseModel
	GroupName string `gorm:"varchar(64);index"`
	Label     string `gorm:"varchar(64)"`
	Value     string `gorm:"varchar(64)"`
}

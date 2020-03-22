package model

import "github.com/jinzhu/gorm"

type PushMsg struct {
	gorm.Model
	BaseModel
	Detail string `gorm:"type:varchar(256)" validate:"required" json:"detail"`
	Status int    `gorm:"type:int(3)" validate:"required" json:"status"`
	SubId  int    `gorm:"type:bigint" validate:"required" json:"subId"`
	UserId uint   `gorm:"type:bigint" validate:"required" json:"userId"`
}

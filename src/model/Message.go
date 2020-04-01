package model

import "github.com/jinzhu/gorm"

type Message struct {
	gorm.Model
	BaseModel
	Title  string `gorm:"type:varchar(64); index" validate:"required" json:"title"`
	Label  string `gorm:"type:varchar(128)" validate:"required" json:"label"`
	Status int    `gorm:"type:int(3);default:0" json:"status"`
	Detail string `gorm:"type:varchar(256)" validate:"required" json:"detail"`
	UserId uint   `gorm:"type:bigint; index" validate:"required" json:"user_id"`
}

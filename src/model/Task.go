package model

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	BaseModel
	Title     string `gorm:"type:varchar(64); index" validate:"required" json:"title"`
	Label     string `gorm:"type:varchar(128)" validate:"required" json:"label"`
	Detail    string `gorm:"type:varchar(256)" validate:"required" json:"detail"`
	Status    int    `gorm:"type:int(3)" validate:"required" json:"status"`
	AlarmTime uint64 `gorm:"type:bigint" validate:"required" json:"alarmTime"`
	UserId    uint   `gorm:"type:bigint; index" validate:"required" json:"userId"`
}

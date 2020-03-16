package model

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Title     string `gorm:"type:varchar(64); index" validate:"required" json:"title"`
	Label     string `gorm:"type:varchar(128)" validate:"required" json:"label"`
	Detail    string `gorm:"type:varchar(256)" validate:"required" json:"detail"`
	Status    uint   `gorm:"type:int(3)" validate:"required" json:"status"`
	AlarmTime uint64 `gorm:"type:bigint" validate:"required" json:"alarmTime"`
	UserID    uint   `gorm:"index" json:"user_id" binding:"required"`
}

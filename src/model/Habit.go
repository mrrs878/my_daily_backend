package model

import "github.com/jinzhu/gorm"

type Habit struct {
	gorm.Model
	BaseModel
	Title      string `gorm:"type:varchar(64); index" validate:"required" json:"title"`
	Label      string `gorm:"type:varchar(128)" json:"label"`
	Status     int    `gorm:"type:int(3)" validate:"required" json:"status"`
	Detail     string `gorm:"type:varchar(256)" validate:"required" json:"detail"`
	SuccessCnt uint   `gorm:"type:bigint" validate:"required" json:"successCnt"`
	FailedCnt  uint   `gorm:"type:bigint" validate:"required" json:"failedCnt"`
	AlarmTime  string `gorm:"type:varchar(16)" validate:"required" json:"alarmTime"`
	AlarmDate  string `gorm:"type:varchar(32)" validate:"required" json:"alarmDate"`
	UserId     uint   `gorm:"type:bigint; index" validate:"required" json:"userId"`
}

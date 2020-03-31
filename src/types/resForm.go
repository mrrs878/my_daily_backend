package types

import (
	"github.com/jinzhu/gorm"
)

type UserInfoForm struct {
	gorm.Model
	Name  string `json:"name"`
	Role  uint   `json:"role"`
	Token string `json:"token"`
}

type TaskInfoForm struct {
	gorm.Model
	Title     string `json:"title"`
	Label     string `json:"title"`
	Detail    string `json:"detail"`
	UserId    uint   `json:"userId"`
	AlarmTime uint64 `json:"alarmTime"`
	Status    int    `json:"status"`
}

type HabitInfoForm struct {
	gorm.Model
	Title      string `json:"title"`
	Label      string `json:"title"`
	Detail     string `json:"detail"`
	Status     int    `json:"status"`
	UserId     uint   `json:"userId"`
	AlarmTime  string `json:"alarmTime"`
	AlarmDate  string `json:"alarmDate"`
	SuccessCnt uint   `json:"successCnt"`
	FailedCnt  uint   `json:"failedCnt"`
}

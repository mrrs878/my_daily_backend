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

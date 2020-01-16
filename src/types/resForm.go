package types

import (
	"demo_1/src/model"
	"github.com/jinzhu/gorm"
)

type UserInfoForm struct {
	gorm.Model
	Name   string        `json:"name"`
	Emails []model.Email `json:"emails"`
	Role   uint          `json:"role"`
	Token  string        `json:"token"`
}

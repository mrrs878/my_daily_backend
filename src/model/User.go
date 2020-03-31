package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name         string `gorm:"type:varchar(64); unique_index" validate:"required" json:"name"`
	PasswordHash string `gorm:"type:varchar(64)" validate:"required"`
	AccessToken  string `gorm:"type:varchar(64)" json:"access_token"`
	Role         uint   `gorm:"type:int(3); default:1" json:"role"`
	Deleted      bool   `gorm:"type:tinyint(1); default:0" json:"deleted"`
	WsId         string `gorm:"type:varchar(128)" json:"ws_id"`
	WsGroup      string `gorm:"type:varchar(128)" json:"ws_group"`
	Tasks        []Task `json:"tasks"`
}

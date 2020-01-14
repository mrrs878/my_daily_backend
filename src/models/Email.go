package models

import "github.com/jinzhu/gorm"

type Email struct {
	gorm.Model
	UserID     int    `gorm:"index" json:"user_id"`
	Email      string `gorm:"type:varchar(100);unique_index" binding:"required" json:"email"`
	UserName   string `gorm:"type:varchar(64)" binding:"required" json:"user_name"`
	Subscribed bool   `gorm:"default:'false'" json:"subscribed"`
}

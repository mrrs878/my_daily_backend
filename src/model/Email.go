package model

import "github.com/jinzhu/gorm"

type Email struct {
	gorm.Model
	UserID     int    `gorm:"index" json:"user_id" binding:"required"`
	Email      string `gorm:"type:varchar(100)" binding:"required" json:"email"`
	Subscribed bool   `gorm:"default:false" json:"subscribed"`
}

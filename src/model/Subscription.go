package model

import "github.com/jinzhu/gorm"

type Subscription struct {
	gorm.Model
	BaseModel
	Endpoint             string `gorm:"type:varchar(511); unique_index" validate:"required" json:"endpoint"`
	ExpirationTime       uint   `gorm:"type:bigint" json:"expirationTime"`
	ApplicationServerKey string `gorm:"type:varchar(511)" json:"applicationServerKey"`
	Auth                 string `gorm:"type:varchar(256)" json:"auth"`
	UserId               uint   `gorm:"type:bigint;index" validate:"required" json:"userId"`
}

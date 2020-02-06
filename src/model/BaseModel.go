package model

type BaseModel struct {
	CreateId uint `gorm:"int(10)"`
	UpdateId uint `gorm:"int(10)"`
	DeleteId uint `gorm:"int(10)"`
}

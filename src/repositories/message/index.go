package message

import (
	"demo_1/src/database"
	"demo_1/src/model"
)

func Add(msg *model.Message) error {
	result := database.DB.Create(&msg)
	return result.Error
}

func Del(msg *model.Message) error {
	result := database.DB.Delete(&msg)
	return result.Error
}

func Update(msg *model.Message) error {
	result := database.DB.Model(msg).Update(msg).Find(msg)
	return result.Error
}

func Index(msg *model.Message) error {
	result := database.DB.Where(&msg).Find(&msg)
	return result.Error
}

func View(msg *[]model.Message) error {
	result := database.DB.Find(msg)
	return result.Error
}

func ViewWithCondition(tasks *[]model.Message, condition interface{}, args ...interface{}) error {
	result := database.DB.Where(condition, args).Find(tasks)
	return result.Error
}

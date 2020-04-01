package message

import (
	"demo_1/src/database"
	"demo_1/src/model"
	"demo_1/src/types"
)

func Add(msg *model.Message) error {
	result := database.DB.Create(&msg)
	return result.Error
}

func Del(msg *model.Message) error {
	err := Update(msg, "id = ?", msg.ID)
	if err != nil {
		return err
	}
	result := database.DB.Delete(&msg)
	return result.Error
}

func Update(newVal *model.Message, condition interface{}, args ...interface{}) error {
	result := database.DB.Table("message").Where(condition, args).Update(newVal).Find(newVal)
	return result.Error
}

func Updates(newVal *types.UpdateNewVal, msg *[]model.Message, condition interface{}, args ...interface{}) error {
	result := database.DB.Table("message").Where(condition, args).Updates(newVal).Find(msg)
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
	result := database.DB.Table("message").Where(condition, args).Find(tasks)
	return result.Error
}

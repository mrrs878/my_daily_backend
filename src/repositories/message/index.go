package message

import (
	"demo_1/src/database"
	"demo_1/src/model"
	"demo_1/src/types"
)

func Add(msg *model.Message) error {
	result := database.DB.Create(&msg).Find(msg)
	return result.Error
}

func Del(msg *model.Message) error {
	_newVal := types.UpdateNewVal{"UpdateId": msg.UpdateId}
	err := Update(&_newVal, msg, "id = ?", msg.ID)
	if err != nil {
		return err
	}
	result := database.DB.Delete(&msg)
	return result.Error
}

func Update(newVal *types.UpdateNewVal, out interface{}, condition interface{}, args ...interface{}) error {
	result := database.DB.Table("message").Where(condition, args).Update(*newVal).Find(out)
	return result.Error
}

func Index(msg *model.Message) error {
	result := database.DB.Where(&msg).Find(&msg)
	return result.Error
}

func View(out *[]model.Message, condition interface{}, args ...interface{}) error {
	result := database.DB.Where(condition, args).Find(out)
	return result.Error
}

func ViewWithCondition(tasks *[]model.Message, condition interface{}, args ...interface{}) error {
	result := database.DB.Table("message").Where(condition, args).Find(tasks)
	return result.Error
}

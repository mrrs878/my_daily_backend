package task

import (
	"demo_1/src/database"
	"demo_1/src/model"
)

func Add(task *model.Task) error {
	result := database.DB.Create(&task)
	return result.Error
}

func Del(task *model.Task) error {
	result := database.DB.Delete(&task)
	return result.Error
}

func Update(task *model.Task) error {
	result := database.DB.Update(&task)
	return result.Error
}

func Index(task *model.Task) error {
	result := database.DB.Where(&task).Find(&task)
	return result.Error
}

func View(task *[]model.Task) error {
	result := database.DB.Find(task)
	return result.Error
}

func ViewWithCondition(tasks *[]model.Task, condition interface{}, args ...interface{}) error {
	result := database.DB.Where(condition, args).Find(tasks)
	return result.Error
}

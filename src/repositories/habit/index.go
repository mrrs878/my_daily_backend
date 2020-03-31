package habit

import (
	"demo_1/src/database"
	"demo_1/src/model"
)

func Add(habit *model.Habit) error {
	result := database.DB.Create(&habit)
	return result.Error
}

func Del(habit *model.Habit) error {
	result := database.DB.Delete(&habit)
	return result.Error
}

func Update(habit *model.Habit) error {
	result := database.DB.Model(habit).Update(habit).Find(habit)
	return result.Error
}

func Index(habit *model.Habit) error {
	result := database.DB.Where(&habit).Find(&habit)
	return result.Error
}

func View(habit *[]model.Habit) error {
	result := database.DB.Find(habit)
	return result.Error
}

func ViewWithCondition(tasks *[]model.Habit, condition interface{}, args ...interface{}) error {
	result := database.DB.Where(condition, args).Find(tasks)
	return result.Error
}

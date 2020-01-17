package user

import (
	"demo_1/src/database"
	"demo_1/src/model"
)

func Add(user *model.User) (interface{}, error) {
	result := database.DB.Create(&user)
	return result.Value, result.Error
}

func Index(user *model.User) error {
	result := database.DB.Where(&user).Find(&user)
	//database.DB.Model(&user).Related(&user.Emails)
	return result.Error
}

func View() {}

func Delete(user *model.User) error {
	result := database.DB.Delete(&user)
	return result.Error
}

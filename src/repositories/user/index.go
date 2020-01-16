package user

import (
	"demo_1/src/database"
	"demo_1/src/model"
)

func Add(user *model.User) (interface{}, error) {
	result := database.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return result.Value, nil
}

func Index(user *model.User) error {
	result := database.DB.Where(&user).Find(&user)
	database.DB.Model(&user).Related(&user.Emails)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func View() {}

func Delete() {}

package dataDict

import (
	"demo_1/src/database"
	"demo_1/src/model"
)

func Add(dataDict *model.DataDict) error {
	result := database.DB.Create(dataDict)
	return result.Error
}

func Delete(dataDict *model.DataDict) error {
	result := database.DB.Delete(dataDict)
	return result.Error
}

func Update(dataDict *model.DataDict) error {
	result := database.DB.Model(&dataDict).Update(dataDict)
	return result.Error
}

func Index(dataDict *model.DataDict) error {
	result := database.DB.Where(dataDict).Find(dataDict)
	return result.Error
}

func View(dataDict *[]model.DataDict) error {
	result := database.DB.Find(dataDict)
	return result.Error
}

func ViewWithCondition(condition *model.DataDict, dataDict *[]model.DataDict) error {
	result := database.DB.Where(condition).Find(dataDict)
	return result.Error
}

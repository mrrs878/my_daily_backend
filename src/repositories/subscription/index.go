package subscription

import (
	"demo_1/src/database"
	"demo_1/src/model"
)

func Add(sub *model.Subscription) (interface{}, error) {
	result := database.DB.Create(&sub)
	return result.Value, result.Error
}

func Index(sub *model.Subscription, condition interface{}, args ...interface{}) error {
	result := database.DB.Where(condition, args).Find(sub)
	return result.Error
}

func ViewWithCondition(subs *[]model.Subscription, condition interface{}, args ...interface{}) error {
	result := database.DB.Where(condition, args).Find(subs)
	return result.Error
}

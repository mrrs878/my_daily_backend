package pushMsg

import (
	"demo_1/src/database"
	"demo_1/src/model"
)

func Add(msg *model.PushMsg) (interface{}, error) {
	result := database.DB.Create(&msg)
	return result.Value, result.Error
}

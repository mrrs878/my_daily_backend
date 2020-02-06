package goods

import (
	"demo_1/src/config"
	"demo_1/src/database"
	"demo_1/src/model"
)

func Add(goods *model.Goods) error {
	result := database.DB.Create(&goods)
	return result.Error
}

func Delete(goods *model.Goods) error {
	result := database.DB.Delete(goods)
	return result.Error
}

func Update(goods *model.Goods) error {
	result := database.DB.Model(goods).Update(goods)
	return result.Error
}

func Index(goods *model.Goods) error {
	result := database.DB.Where(goods)
	return result.Error
}

func ViewByClassAndPage(class uint, page uint, goods *[]model.Goods) error {
	result := database.DB.Offset((page - 1) * config.QueryPageSize).Where(&model.Goods{Class: class}).Find(&goods)
	return result.Error
}

package database

import (
	"demo_1/src/config"
	"demo_1/src/model"
	"fmt"
	"github.com/jinzhu/gorm"
)
import _ "github.com/jinzhu/gorm/dialects/mysql"

var DB *gorm.DB

func SetUpDatabase() {
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DatabaseUsername,
		config.DatabasePassword,
		config.DatabaseURl,
		config.DatabasePort,
		config.DatabaseName)
	db, err := gorm.Open("mysql", dsn)
	db.LogMode(true)
	if err != nil {
		panic(err)
		return
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.Email{}).
		AutoMigrate(&model.User{})

	DB = db
	fmt.Println("|    connected to database        |")
	//defer db.Close()
}

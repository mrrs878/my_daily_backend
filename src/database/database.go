package database

import (
	"demo_1/src/models"
	"fmt"
	"github.com/jinzhu/gorm"
)
import _ "github.com/jinzhu/gorm/dialects/mysql"

var DB *gorm.DB

func SetUpDatabase() {
	db, err := gorm.Open("mysql", "e_market:Shkji5CMzY347XBM@tcp(www.sumcet.com:3306)/e_market?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	if err != nil {
		panic(err)
		return
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&models.Email{})
	DB = db
	fmt.Println("|    connected to database        |")
	//defer db.Close()
}

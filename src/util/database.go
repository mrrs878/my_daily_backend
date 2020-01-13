package util

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/mysql"

type Student struct {
	Id    int
	Name  string
	Age   int
	Sex   byte
	Phone string
}

func SetUpDatabase() {
	db, err := gorm.Open("mysql", "e_market:Shkji5CMzY347XBM@tcp(www.sumcet.com:3306)/e_market?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	defer db.Close()
}

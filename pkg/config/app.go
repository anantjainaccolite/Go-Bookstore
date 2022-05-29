package config

// this file connects the project withDatabase

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//this would be returned from the file
var (
	db *gorm.DB
)

func Connect() { //to open a  connection with our database
	dsn := "username:password@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

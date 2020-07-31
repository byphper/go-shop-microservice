package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var Db *gorm.DB

func init() {
	Db = GetDb()
}

func GetDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:ayong@/user_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(true)
	db.DB().SetMaxOpenConns(20)
	db.DB().SetMaxIdleConns(5)
	return db
}

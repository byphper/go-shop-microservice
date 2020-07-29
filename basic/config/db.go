package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func GetDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:ayong@/xxswkl?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(true)
	db.DB().SetMaxOpenConns(20)
	db.DB().SetMaxIdleConns(5)
	return db
}

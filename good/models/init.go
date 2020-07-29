package models

import (
	"github.com/jinzhu/gorm"
	"go-shop-microservice/basic/config"
)

var Db *gorm.DB

func init() {
	Db = config.GetDb()
}

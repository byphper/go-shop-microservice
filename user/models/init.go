package models

import (
	"github.com/jinzhu/gorm"
	"xxswkl/basic/config"
)

var Db *gorm.DB

func init() {
	Db = config.GetDb()
}

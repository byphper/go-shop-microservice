package models

import (
	"github.com/jinzhu/gorm"
)

type AttrValue struct {
	gorm.Model
	AttrID uint
	Value  string `gorm:"type:varchar(10)"`
	Desc   string `gorm:"type:varchar(50)"`
}

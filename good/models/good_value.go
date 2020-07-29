package models

import (
	"github.com/jinzhu/gorm"
)

type GoodAttrValue struct {
	gorm.Model
	GoodAttrId uint
	Value       string `gorm:"type:varchar(10)"`
	Desc       string `gorm:"type:varchar(50)"`
}

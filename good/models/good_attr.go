package models

import (
	"github.com/jinzhu/gorm"
)

type GoodAttr struct {
	gorm.Model
	Name           string `gorm:"type:varchar(10);unique_index"`
	Desc           string `gorm:"type:varchar(50)"`
	GoodAttrValues []GoodAttrValue
}

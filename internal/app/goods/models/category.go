package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name    string `gorm:"type:varchar(10);unique_index"`
	Desc    string `gorm:"type:varchar(50)"`
	PId     uint   `gorm:"index"`
	Path    string `gorm:"type:varchar(255)"`
	LogoUrl string `gorm:"type:varchar(255)"`
	Status  uint   `gorm:"type:tinyint(1);DEFAULT '1'"`
}

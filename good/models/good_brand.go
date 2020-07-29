package models

import (
	"github.com/jinzhu/gorm"
)

type GoodBrand struct {
	gorm.Model
	Name    string `gorm:"type:varchar(10);"`
	Desc    string `gorm:"type:varchar(100)"`
	LogoUrl string `gorm:"type:varchar(255)"`
	Status  uint   `gorm:"type:tinyint(1)"`
}

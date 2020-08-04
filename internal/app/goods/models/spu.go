package models

import (
	"github.com/jinzhu/gorm"
)

type Spu struct {
	gorm.Model
	Name        string `gorm:"type:varchar(50);"`
	BrandId     uint   `gorm:"index"`
	CategoryId  uint   `gorm:"index"`
	Brand       Brand
	Category    Category
	Desc        string `gorm:"type:varchar(255)"`
	LogoUrl     string `gorm:"type:varchar(255)"`
	MainUrl     string `gorm:"type:varchar(255)"`
	BannerUrl   string `gorm:"type:varchar(255)"`
	SellPrice   uint
	MarketPrice uint
	sales       uint
	Unit        string `gorm:"type:varchar(10);"`
	Status      uint   `gorm:"type:tinyint(1)"`
}

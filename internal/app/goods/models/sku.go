package models

import (
	"github.com/jinzhu/gorm"
)

type Sku struct {
	gorm.Model
	SpuId       uint `gorm:"index"`
	Spu         Spu
	Name        string `gorm:"type:varchar(50);"`
	Desc        string `gorm:"type:varchar(255)"`
	LogoUrl     string `gorm:"type:varchar(255)"`
	MainUrl     string `gorm:"type:varchar(255)"`
	BannerUrl   string `gorm:"type:varchar(255)"`
	SellPrice   uint
	MarketPrice uint
	Stock       uint
	Status      uint `gorm:"type:tinyint(1)"`
}

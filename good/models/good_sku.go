package models

import (
	"github.com/jinzhu/gorm"
)

type GoodSku struct {
	gorm.Model
	GoodSpuId       uint
	GoodSpu         GoodSpu
	Name            string `gorm:"type:varchar(50);"`
	Desc            string `gorm:"type:varchar(255)"`
	LogoUrl         string `gorm:"type:varchar(255)"`
	MainUrl         string `gorm:"type:varchar(255)"`
	BannerUrl       string `gorm:"type:varchar(255)"`
	GoodsAttrValues string `gorm:"type:varchar(255)"`
	SellPrice       uint
	MarketPrice     uint
	Quantity        uint
	Status          uint `gorm:"type:tinyint(1)"`
}

package models

import "github.com/jinzhu/gorm"

type GoodSpuSkuAttrMap struct {
	gorm.Model
	GoodSpuId         uint `gorm:"index:good_spu_id"`
	GoodSkuId         uint
	GoodAttrId        uint
	GoodAttrName      string `gorm:"type:varchar(10);"`
	GoodAttrValueId   uint
	GoodAttrValueName string `gorm:"type:varchar(10);"`
}

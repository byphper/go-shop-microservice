package models

import (
	"github.com/jinzhu/gorm"
)

type SpuSkuAttrMap struct {
	gorm.Model
	SpuId       uint `gorm:"index:good_spu_id"`
	SkuId       uint
	AttrId      uint
	Attr        string `gorm:"type:varchar(50);"`
	AttrValueId uint
	AttrValue   string `gorm:"type:varchar(50);"`
}

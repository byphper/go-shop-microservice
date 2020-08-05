package repository

import (
	"github.com/pkg/errors"
	"go-shop-microservice/internal/app/goods/entities"
	"go-shop-microservice/internal/app/goods/models"
)

type SkuResp struct {
}

func (skuResp *SkuResp) Create(attrs []map[uint]uint, name, desc, logoUrl, mainUrl, BannelUrl string, spuId, stock, sellPrice, marketPrice uint) (sku entities.Sku, err error) {
	spuResp := &SpuResp{}
	spu, err := spuResp.Get(spuId)
	if err != nil {
		return sku, errors.Wrap(err, "无效的SPU")
	}
	model := &models.Sku{
		Name:        name,
		Desc:        desc,
		SpuId:       spu.Id,
		LogoUrl:     logoUrl,
		MainUrl:     mainUrl,
		BannerUrl:   BannelUrl,
		SellPrice:   sellPrice,
		MarketPrice: marketPrice,
		Stock:       stock,
	}
	if err = models.Db.Create(model).Error; err != nil {
		return
	}
	skuResp.composeField(&sku, model)
	return
}

func (skuResp *SkuResp) Update(sku entities.Sku) (err error) {
	model := models.Sku{
		Name:        sku.Name,
		Desc:        sku.Desc,
		LogoUrl:     sku.LogoUrl,
		MainUrl:     sku.MainUrl,
		BannerUrl:   sku.BannerUrl,
		SellPrice:   sku.SellPrice,
		MarketPrice: sku.MarketPrice,
		Stock:       sku.Stock,
	}
	err = models.Db.Model(&model).Where("id =?", sku.Id).Updates(model).Error
	return
}

func (skuResp *SkuResp) Get(id uint) (sku entities.Sku, err error) {
	model := &models.Sku{}
	if err = models.Db.Where("id =?", id).First(model).Error; err != nil {
		return
	}
	spuResp := &SpuResp{}
	spu, err := spuResp.Get(model.SpuId)
	if err != nil {
		return
	}
	sku.Spu = spu
	skuResp.composeField(&sku, model)
	return
}

func (skuResp *SkuResp) composeField(sku *entities.Sku, model *models.Sku) {
	sku.Id = model.ID
	sku.Name = model.Name
	sku.Desc = model.Desc
	sku.LogoUrl = model.LogoUrl
	sku.MainUrl = model.MainUrl
	sku.BannerUrl = model.BannerUrl
	sku.SellPrice = model.SellPrice
	sku.MarketPrice = model.MarketPrice
	sku.Stock = model.Stock
	sku.Status = model.Status
	sku.CreatedAt = model.CreatedAt
}

package services

import (
	"go-shop-microservice/internal/app/goods/entities"
	"go-shop-microservice/internal/app/goods/repository"
)

type Sku struct {
	SkuResp *repository.SkuResp
}

func (skuService *Sku) Create(name, desc, attrValues, logoUrl, mainUrl, BannelUrl string, spuId, stock, sellPrice, marketPrice uint) (entities.Sku, error) {
	return skuService.SkuResp.Create(name, desc, attrValues, logoUrl, mainUrl, BannelUrl, spuId, stock, sellPrice, marketPrice)
}

func (skuService *Sku) Get(skuId uint) (entities.Sku, error) {
	return skuService.SkuResp.Get(skuId)
}

func (skuService *Sku) Update(sku entities.Sku) error {
	return skuService.SkuResp.Update(sku)
}

func (skuService *Sku) GetSkuByAttrValues(spuId uint, attrValues string) (sku entities.Sku, err error) {
	return skuService.SkuResp.GetSkuByAttrValues(spuId, attrValues)
}

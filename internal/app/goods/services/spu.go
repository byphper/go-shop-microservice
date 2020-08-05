package services

import (
	"go-shop-microservice/internal/app/goods/entities"
	"go-shop-microservice/internal/app/goods/repository"
)

type Spu struct {
	SpuResp *repository.SpuResp
}

func (skuService *Spu) Create(name, desc, logoUrl, mainUrl, BannelUrl, unit, attrIds string, cid, bid, sellPrice, marketPrice uint) (entities.Spu, error) {
	return skuService.SpuResp.Create(name, desc, logoUrl, mainUrl, BannelUrl, unit, attrIds, cid, bid, sellPrice, marketPrice)
}

func (skuService *Spu) Get(spuId uint) (entities.Spu, error) {
	return skuService.SpuResp.Get(spuId)
}

func (skuService *Spu) Update(spu entities.Spu) error {
	return skuService.SpuResp.Update(spu)
}

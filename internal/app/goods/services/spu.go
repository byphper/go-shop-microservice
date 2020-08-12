package services

import (
	"go-shop-microservice/internal/app/goods/entities"
	"go-shop-microservice/internal/app/goods/repository"
)

type Spu struct {
	SpuResp *repository.SpuResp
}

func (spuService *Spu) Create(name, desc, logoUrl, mainUrl, BannelUrl, unit, attrIds string, cid, bid, sellPrice, marketPrice uint) (entities.Spu, error) {
	return spuService.SpuResp.Create(name, desc, logoUrl, mainUrl, BannelUrl, unit, attrIds, cid, bid, sellPrice, marketPrice)
}

func (spuService *Spu) Get(spuId uint) (entities.Spu, error) {
	return spuService.SpuResp.Get(spuId)
}

func (spuService *Spu) Update(spu entities.Spu) error {
	return spuService.SpuResp.Update(spu)
}

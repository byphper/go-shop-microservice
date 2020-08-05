package services

import (
	"go-shop-microservice/internal/app/goods/repository"
	"testing"
)

func TestSkuService_Create(t *testing.T) {
	skuService := &Sku{SkuResp: &repository.SkuResp{}}
	name, desc, logoUrl, mainUrl, BannelUrl := "美特斯邦威", "不错", "", "", ""
	spuId, stock, sp, mp := uint(1), uint(999), uint(1000), uint(1200)
	attr, err := skuService.Create(name, desc, logoUrl, mainUrl, BannelUrl, spuId, stock, sp, mp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(attr)
}

func TestSkuService_Get(t *testing.T) {
	SkuService := &Sku{SkuResp: &repository.SkuResp{}}
	var SkuId uint = 1
	Sku, err := SkuService.Get(SkuId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(Sku)
}

func TestSkuService_Update(t *testing.T) {
	SkuService := Sku{SkuResp: &repository.SkuResp{}}
	var SkuId uint = 1
	Sku, _ := SkuService.Get(SkuId)
	Sku.Desc = "是改了的"
	err := SkuService.Update(Sku)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(Sku)
}

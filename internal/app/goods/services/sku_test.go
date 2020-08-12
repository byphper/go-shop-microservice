package services

import (
	"go-shop-microservice/internal/app/goods/repository"
	"testing"
)

func TestSkuService_Create(t *testing.T) {
	skuService := &Sku{SkuResp: &repository.SkuResp{}}
	name, desc, logoUrl, mainUrl, BannelUrl := "美", "不错", "", "", ""
	spuId, stock, sp, mp := uint(1), uint(999), uint(1000), uint(1200)
	attrValues := "[{1:1},{2:6}]"
	attr, err := skuService.Create(name, desc, attrValues, logoUrl, mainUrl, BannelUrl, spuId, stock, sp, mp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(attr)
}

func TestSkuService_Get(t *testing.T) {
	skuService := &Sku{SkuResp: &repository.SkuResp{}}
	var SkuId uint = 1
	sku, err := skuService.Get(SkuId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sku)
}

func TestSkuService_GetSkuByAttrValue(t *testing.T) {
	skuService := &Sku{SkuResp: &repository.SkuResp{}}
	var spuId uint = 1
	var attrValues = "[{\"attr_id\":1,\"value_id\":2},{\"attr_id\":2,\"value_id\":6}]"
	sku, err := skuService.GetSkuByAttrValues(spuId, attrValues)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sku)
}

func TestSkuService_Update(t *testing.T) {
	skuService := Sku{SkuResp: &repository.SkuResp{}}
	var SkuId uint = 1
	sku, _ := skuService.Get(SkuId)
	sku.Desc = "是改了的"
	err := skuService.Update(sku)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sku)
}

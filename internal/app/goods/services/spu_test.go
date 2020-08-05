package services

import (
	"go-shop-microservice/internal/app/goods/repository"
	"testing"
)

func TestSpuService_Create(t *testing.T) {
	spuService := &Spu{SpuResp: &repository.SpuResp{}}
	name, desc, logoUrl, mainUrl, BannelUrl, unit := "Vivaia男鞋", "好看", "", "", "", "双"
	attrIds := "1-2"
	cid, bid, sp, mp := uint(1), uint(1), uint(1000), uint(1200)
	attr, err := spuService.Create(name, desc, logoUrl, mainUrl, BannelUrl, unit, attrIds, cid, bid, sp, mp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(attr)
}

func TestSpuService_Get(t *testing.T) {
	spuService := &Spu{SpuResp: &repository.SpuResp{}}
	var spuId uint = 1
	spu, err := spuService.Get(spuId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(spu)
}

func TestSpuService_Update(t *testing.T) {
	spuService := Spu{SpuResp: &repository.SpuResp{}}
	var spuId uint = 1
	spu, _ := spuService.Get(spuId)
	spu.Desc = "是改了的"
	err := spuService.Update(spu)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(spu)
}

package services

import (
	"go-shop-microservice/internal/app/goods/repository"
	"testing"
)

func TestBrandService_CreateBrand(t *testing.T) {
	brandService := Brand{BrandResp: &repository.BrandResp{}}
	name, desc,logoUrl := "阿迪", "niubi",""
	attr, err := brandService.Create(name, desc,logoUrl)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(attr)
}

func TestBrandService_Get(t *testing.T) {
	brandService := Brand{BrandResp: &repository.BrandResp{}}
	var brandId uint = 1
	brand, err := brandService.Get(brandId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(brand)
}

func TestBrandService_Update(t *testing.T) {
	brandService := Brand{BrandResp: &repository.BrandResp{}}
	var brandId uint = 1
	attr, _ := brandService.Get(brandId)
	attr.Desc = "我改了这个描述"
	err := brandService.Update(attr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(attr)
}

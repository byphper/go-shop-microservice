package services

import (
	"go-shop-microservice/internal/app/goods/entities"
	"go-shop-microservice/internal/app/goods/repository"
)

type Brand struct {
	BrandResp *repository.BrandResp
}

func (cateService *Brand) Create(name string, desc string, logoUrl string) (entities.Brand, error) {
	return cateService.BrandResp.Create(name, desc, logoUrl)
}

func (cateService *Brand) Get(brandId uint) (entities.Brand, error) {
	return cateService.BrandResp.Get(brandId)
}

func (cateService *Brand) Update(brand entities.Brand) error {
	return cateService.BrandResp.Update(brand)
}

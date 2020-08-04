package repository

import (
	"go-shop-microservice/internal/app/goods/entities"
	"go-shop-microservice/internal/app/goods/models"
)

type BrandResp struct {
}

func (brandResp *BrandResp) Create(name string, desc string, logoUrl string) (brand entities.Brand, err error) {
	model := &models.Brand{
		Name:    name,
		Desc:    desc,
		LogoUrl: logoUrl,
	}
	if err = models.Db.Create(model).Error; err != nil {
		return
	}
	brand.Id = model.ID
	brand.Name = model.Name
	brand.Desc = model.Desc
	brand.LogoUrl = model.LogoUrl
	return
}

func (brandResp *BrandResp) Update(brand entities.Brand) (err error) {
	model := &models.Brand{
		Name:    brand.Name,
		Desc:    brand.Desc,
		LogoUrl: brand.LogoUrl,
	}
	err = models.Db.Model(&model).Where("id =?", brand.Id).Updates(model).Error
	return
}

func (brandResp *BrandResp) Get(id uint) (brand entities.Brand, err error) {
	model := &models.Brand{}
	if err = models.Db.Where("id =?", id).First(model).Error; err != nil {
		return
	}
	brand.Id = model.ID
	brand.Name = model.Name
	brand.Desc = model.Desc
	brand.LogoUrl = model.LogoUrl
	return
}

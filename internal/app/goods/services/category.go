package services

import (
	"go-shop-microservice/internal/app/goods/entities"
	"go-shop-microservice/internal/app/goods/repository"
)

type Category struct {
	CateResp *repository.CategoryResp
}

func (cateService *Category) Create(name string, desc string, logoUrl string, pid uint) (entities.Category, error) {
	return cateService.CateResp.Create(name, desc, logoUrl, pid)
}

func (cateService *Category) Get(cateId uint) (entities.Category, error) {
	return cateService.CateResp.Get(cateId)
}

func (cateService *Category) GetAllChilds(cateId uint) ([]entities.Category, error) {
	return cateService.CateResp.GetAllChilds(cateId)
}

func (cateService *Category) GetDirectChilds(cateId uint) ([]entities.Category, error) {
	return cateService.CateResp.GetDirectChilds(cateId)
}

func (cateService *Category) Update(cate entities.Category) error {
	return cateService.CateResp.Update(cate)
}

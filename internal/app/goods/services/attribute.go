package services

import (
	"go-shop-microservice/internal/app/goods/entities"
	"go-shop-microservice/internal/app/goods/repository"
)

type Attribute struct {
	AttrResp *repository.AttributeResp
}

func (attributeService *Attribute) Create(name string, desc string) (entities.Attribute, error) {
	return attributeService.AttrResp.Create(name, desc)
}

func (attributeService *Attribute) CreateValue(attrId uint, name string, desc string) error {
	return attributeService.AttrResp.CreateValues(attrId, name, desc)
}

func (attributeService *Attribute) Get(attrId uint) (entities.Attribute, error) {
	return attributeService.AttrResp.Get(attrId)
}

func (attributeService *Attribute) Update(attr entities.Attribute) error {
	return attributeService.AttrResp.Update(attr)
}

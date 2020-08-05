package repository

import (
	"go-shop-microservice/internal/app/goods/entities"
	"go-shop-microservice/internal/app/goods/models"
)

type AttributeResp struct {
}

func (attrResp *AttributeResp) Create(name string, desc string) (attr entities.Attribute, err error) {
	model := &models.Attr{
		Name: name,
		Desc: desc,
	}
	if err = models.Db.Create(model).Error; err != nil {
		return
	}
	attr.Id = model.ID
	attr.Name = model.Name
	attr.Desc = model.Desc
	return
}

func (attrResp *AttributeResp) CreateValues(attrId uint, value string, desc string) (err error) {
	model := &models.AttrValue{
		Value:  value,
		Desc:   desc,
		AttrID: attrId,
	}
	if err = models.Db.Create(model).Error; err != nil {
		return
	}
	return nil
}

func (attrResp *AttributeResp) Update(attr entities.Attribute) (err error) {
	model := &models.Attr{
		Name: attr.Name,
		Desc: attr.Desc,
	}
	err = models.Db.Model(&model).Where("id =?", attr.Id).Updates(model).Error
	return
}

func (attrResp *AttributeResp) Get(id uint) (attr entities.Attribute, err error) {
	model := &models.Attr{}
	if err = models.Db.Where("id =?", id).Preload("AttrValues").First(model).Error; err != nil {
		return
	}
	attr.Id = model.ID
	attr.Name = model.Name
	attr.Desc = model.Desc
	if model.AttrValues != nil {
		for _, v := range model.AttrValues {
			attr.Values = append(attr.Values, map[uint]string{v.ID: v.Value})
		}
	}
	return
}

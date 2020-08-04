package services

import (
	"go-shop-microservice/internal/app/goods/repository"
	"testing"
)

func TestAttributeService_CreateAttr(t *testing.T) {
	attrService := Attribute{AttrResp: &repository.AttributeResp{}}
	name, desc := "尺寸", "展示尺寸"
	attr, err := attrService.Create(name, desc)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(attr)
}

func TestAttributeService_CreateAttrValue(t *testing.T) {
	attrService := Attribute{AttrResp: &repository.AttributeResp{}}
	var attrId uint = 2
	name, desc := "M", "middle"
	err := attrService.CreateValue(attrId, name, desc)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("创建成功")
}

func TestAttributeService_Get(t *testing.T) {
	attrService := Attribute{AttrResp: &repository.AttributeResp{}}
	var attrId uint = 2
	attr, err := attrService.Get(attrId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(attr)
}

func TestAttributeService_Update(t *testing.T) {
	attrService := Attribute{AttrResp: &repository.AttributeResp{}}
	var attrId uint = 2
	attr, _ := attrService.Get(attrId)
	attr.Desc = "是改了的"
	err := attrService.Update(attr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(attr)
}

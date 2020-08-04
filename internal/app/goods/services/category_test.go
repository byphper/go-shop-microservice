package services

import (
	"go-shop-microservice/internal/app/goods/repository"
	"testing"
)

func TestCategoryService_CreateCategory(t *testing.T) {
	CategoryService := Category{CateResp: &repository.CategoryResp{}}
	name, desc, logoUrl, pid := "春装男-中年", "cloth", "", uint(5)
	attr, err := CategoryService.Create(name, desc, logoUrl, pid)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(attr)
}

func TestCategoryService_Get(t *testing.T) {
	CategoryService := Category{CateResp: &repository.CategoryResp{}}
	var CategoryId uint = 1
	category, err := CategoryService.Get(CategoryId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(category)
}

func TestCategoryService_GetAllChilds(t *testing.T) {
	CategoryService := Category{CateResp: &repository.CategoryResp{}}
	var CategoryId uint = 2
	categories, err := CategoryService.GetAllChilds(CategoryId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(categories)
}

func TestCategoryService_GetDirectChilds(t *testing.T) {
	CategoryService := Category{CateResp: &repository.CategoryResp{}}
	var CategoryId uint = 1
	categories, err := CategoryService.GetDirectChilds(CategoryId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(categories)
}

func TestCategoryService_Update(t *testing.T) {
	CategoryService := Category{CateResp: &repository.CategoryResp{}}
	var CategoryId uint = 5
	cate, _ := CategoryService.Get(CategoryId)
	cate.Desc = "呵呵"
	cate.PId = 3
	err := CategoryService.Update(cate)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cate)
}

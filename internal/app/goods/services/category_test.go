package services

import (
	"go-shop-microservice/internal/app/goods/repository"
	"testing"
)

func TestCategoryService_CreateCategory(t *testing.T) {
	CategoryService := Category{CateResp: &repository.CategoryResp{}}
	name, desc, logoUrl, pid := "儿童夏装", "cloth", "", uint(4)
	attr, err := CategoryService.Create(name, desc, logoUrl, pid)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(attr)
}

func TestCategoryService_Get(t *testing.T) {
	CategoryService := Category{CateResp: &repository.CategoryResp{}}
	var CategoryId uint = 1
	Category, err := CategoryService.Get(CategoryId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(Category)
}

func TestCategoryService_Update(t *testing.T) {
	CategoryService := Category{CateResp: &repository.CategoryResp{}}
	var CategoryId uint = 5
	cate, _ := CategoryService.Get(CategoryId)
	cate.Desc = "我换PID了"
	cate.PId = 2
	err := CategoryService.Update(cate)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cate)
}

package repository

import (
	"errors"
	"go-shop-microservice/internal/app/goods/entities"
	"go-shop-microservice/internal/app/goods/models"
	"strconv"
	"strings"
	"sync"
)

type CategoryResp struct {
	pathModifylock sync.Mutex
}

func (categoryResp *CategoryResp) Create(name string, desc string, logoUrl string, pid uint) (cate entities.Category, err error) {
	model := &models.Category{
		Name:    name,
		Desc:    desc,
		LogoUrl: logoUrl,
		PId:     pid,
	}
	tx := models.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err = models.Db.Create(model).Error; err != nil {
		tx.Rollback()
		return
	}
	categoryResp.pathModifylock.Lock() //因为可能会并发修改path 所以这里需要做加锁处理
	if err = categoryResp.deepUpdateCategoryPath(model, model.ID); err != nil {
		categoryResp.pathModifylock.Unlock()
		tx.Rollback()
		return
	}
	categoryResp.pathModifylock.Unlock()
	cate.Id = model.ID
	cate.Name = model.Name
	cate.Desc = model.Desc
	cate.LogoUrl = model.LogoUrl
	cate.PId = model.PId
	return
}

//递归更新分类path
func (categoryResp *CategoryResp) deepUpdateCategoryPath(cateModel *models.Category, id uint) error {
	if cateModel.PId > 0 {
		model := &models.Category{}
		if err := models.Db.Where("id =?", cateModel.PId).First(model).Error; err != nil {
			return err
		}
		pathStr := model.Path + strconv.Itoa(int(id)) + "-"
		if err := models.Db.Model(model).Update("path", pathStr).Error; err != nil {
			return err
		}
		if model.PId > 0 {
			return categoryResp.deepUpdateCategoryPath(model, id)
		}
	}
	return nil
}

//递归删除分类path
func (categoryResp *CategoryResp) deepDeleteCategoryPath(cateModel *models.Category, id uint) error {
	if cateModel.PId > 0 {
		model := &models.Category{}
		if err := models.Db.Where("id =?", cateModel.PId).First(model).Error; err != nil {
			return err
		}
		pathSlice := strings.Split(model.Path, "-")
		var index = 0
		for i, v := range pathSlice {
			if v == strconv.Itoa(int(id)) {
				index = i
				continue
			}
		}
		newPathSlice := append(pathSlice[:index], pathSlice[index+1:]...)
		newPathStr := strings.Join(newPathSlice, "-") + "-"
		if err := models.Db.Model(model).Update("path", newPathStr).Error; err != nil {
			return err
		}
		if model.PId > 0 {
			return categoryResp.deepDeleteCategoryPath(model, id)
		}
	}
	return nil
}

func (categoryResp *CategoryResp) Update(cate entities.Category) (err error) {
	model := &models.Category{
		Name:    cate.Name,
		Desc:    cate.Desc,
		LogoUrl: cate.LogoUrl,
		PId:     cate.PId,
	}
	oldCate, err := categoryResp.Get(cate.Id)
	if err != nil {
		return errors.New("无效的分类")
	}
	tx := models.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err = models.Db.Model(&model).Where("id =?", cate.Id).Updates(model).Error
	if err != nil {
		tx.Rollback()
		return errors.New("更新失败")
	}
	//如果修改了pid 则需要递归处理包含了该分类的path部分
	if oldCate.PId != cate.PId {
		categoryResp.pathModifylock.Lock()
		model.PId = oldCate.PId
		err = categoryResp.deepDeleteCategoryPath(model, cate.Id)
		if err != nil {
			categoryResp.pathModifylock.Unlock()
			tx.Rollback()
			return errors.New("删除旧PATH失败")
		}
		model.PId = cate.PId
		err = categoryResp.deepUpdateCategoryPath(model, cate.Id)
		if err != nil {
			categoryResp.pathModifylock.Unlock()
			tx.Rollback()
			return errors.New("更新PATH失败")
		}
		categoryResp.pathModifylock.Unlock()
	}
	tx.Commit()
	return
}

func (categoryResp *CategoryResp) Get(id uint) (cate entities.Category, err error) {
	model := &models.Category{}
	if err = models.Db.Where("id =?", id).First(model).Error; err != nil {
		return
	}
	cate.Id = model.ID
	cate.Name = model.Name
	cate.Desc = model.Desc
	cate.LogoUrl = model.LogoUrl
	cate.PId = model.PId
	return
}

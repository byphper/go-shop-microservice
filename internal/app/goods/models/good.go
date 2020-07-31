package models

import (
	"errors"
	"strconv"
	"strings"
)

const (
	StatusEnable  = 1
	StatusDisable = 0
)

type GoodDaoImpl struct {
}

//创建品牌
func (gd *GoodDaoImpl) CreateBrand(name string, desc string, logoUrl string) (*GoodBrand, error) {
	brand := &GoodBrand{
		Name:    name,
		Desc:    desc,
		LogoUrl: logoUrl,
		Status:  StatusEnable,
	}
	err := Db.Create(brand).Error
	if err != nil {
		return nil, err
	}
	return brand, nil
}

//创建属性
func (gd *GoodDaoImpl) CreateAttr(name string, desc string) (*GoodAttr, error) {
	attr := &GoodAttr{
		Name: name,
		Desc: desc,
	}
	err := Db.Create(attr).Error
	if err != nil {
		return nil, err
	}
	return attr, nil
}

//创建属性值
func (gd *GoodDaoImpl) CreateAttrValue(attrId uint, value string, desc string) (*GoodAttrValue, error) {
	attrValue := &GoodAttrValue{
		Value:      value,
		Desc:       desc,
		GoodAttrId: attrId,
	}
	err := Db.Create(attrValue).Error
	if err != nil {
		return nil, err
	}
	return attrValue, nil
}

//创建分类
func (gd *GoodDaoImpl) CreateCategory(pid uint, name string, desc string, logoUrl string) (*GoodCategory, error) {
	category := &GoodCategory{
		PId:     pid,
		Name:    name,
		Desc:    desc,
		LogoUrl: logoUrl,
		Status:  StatusEnable,
	}
	tx := Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := Db.Create(category).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := gd.DeepUpdateCategoryPath(category, category.ID); err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

//递归更新分类path
func (gd *GoodDaoImpl) DeepUpdateCategoryPath(category *GoodCategory, id uint) error {
	if category.PId > 0 {
		pCategory, err := gd.GetCategoryById(category.PId, "id", "path", "p_id")
		if err != nil {
			return errors.New("无效的PID")
		}
		path := pCategory.Path + strconv.Itoa(int(id)) + "-"
		if err := Db.Model(pCategory).Update("path", path).Error; err != nil {
			return err
		}
		if pCategory.PId > 0 {
			return gd.DeepUpdateCategoryPath(&pCategory, id)
		}
	}
	return nil
}

//获取一个分类信息
func (gd *GoodDaoImpl) GetCategoryById(id uint, fields ...string) (GoodCategory, error) {
	category := GoodCategory{}
	if fields == nil {
		fields = []string{"*"}
	}
	if err := Db.Select(fields).Where("id =?", id).First(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

//创建spu
func (gd *GoodDaoImpl) CreateSpu(spu *GoodSpu) (*GoodSpu, error) {
	err := Db.Create(spu).Error
	if err != nil {
		return nil, err
	}
	return spu, nil
}

//更新spu
func (gd *GoodDaoImpl) UpdateSpu(spu *GoodSpu) (*GoodSpu, error) {
	err := Db.Model(spu).Update(spu).Error
	if err != nil {
		return nil, err
	}
	return spu, nil
}

//更新sku
func (gd *GoodDaoImpl) UpdateSku(sku *GoodSku) (*GoodSku, error) {
	err := Db.Model(sku).Update(sku).Error
	if err != nil {
		return nil, err
	}
	return sku, nil
}

//获取spu信息
func (gd *GoodDaoImpl) GetSpuById(id uint, fields ...string) (*GoodSpu, error) {
	spu := &GoodSpu{}
	if fields == nil {
		fields = []string{"*"}
	}
	if err := Db.Select(fields).Where("id =?", id).First(spu).Error; err != nil {
		return nil, err
	}
	if err := Db.Model(spu).Related(&spu.GoodBrand).Related(&spu.GoodCategory).Error; err != nil {
		return nil, err
	}
	return spu, nil
}

//获取sku信息
func (gd *GoodDaoImpl) GetSkuById(id uint, fields ...string) (*GoodSku, error) {
	sku := &GoodSku{}
	if fields == nil {
		fields = []string{"*"}
	}
	if err := Db.Select(fields).Where("id =?", id).First(sku).Error; err != nil {
		return nil, err
	}
	return sku, nil
}

//创建sku
func (gd *GoodDaoImpl) CreateSku(sku *GoodSku) (*GoodSku, error) {
	tx := Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err := Db.Create(sku).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if sku.GoodsAttrValues != "" {
		var attrValues []GoodAttrValue
		attarValuesId := strings.Split(sku.GoodsAttrValues, ",")
		if err := Db.Where("id IN (?)", attarValuesId).Find(&attrValues).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		var attrIds = make([]uint, len(attrValues))
		for _, v := range attrValues {
			attrIds = append(attrIds, v.GoodAttrId)
		}
		var attrs []GoodAttr
		if err := Db.Where("id IN (?)", attrIds).Find(&attrs).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		var attrsMap = make(map[uint]string, len(attrs))
		for _, v := range attrs {
			attrsMap[v.ID] = v.Name
		}
		for _, v := range attrValues {
			spuSkuAttrMap := GoodSpuSkuAttrMap{
				GoodSpuId:         sku.GoodSpuId,
				GoodSkuId:         sku.ID,
				GoodAttrId:        v.GoodAttrId,
				GoodAttrName:      attrsMap[v.GoodAttrId],
				GoodAttrValueId:   v.ID,
				GoodAttrValueName: v.Value,
			}
			if err := Db.Create(&spuSkuAttrMap).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return sku, nil
}

package repository

import (
	"encoding/json"
	"github.com/pkg/errors"
	"go-shop-microservice/internal/app/goods/entities"
	"go-shop-microservice/internal/app/goods/models"
	"go-shop-microservice/internal/pkg/util"
	"sort"
)

type SkuResp struct {
}

func (skuResp *SkuResp) Create(name, desc, attrValues, logoUrl, mainUrl, BannelUrl string, spuId, stock, sellPrice, marketPrice uint) (sku entities.Sku, err error) {
	spuResp := new(SpuResp)
	spu, err := spuResp.Get(spuId)
	if err != nil {
		return sku, errors.Wrap(err, "无效的SPU")
	}
	model := &models.Sku{
		Name:        name,
		Desc:        desc,
		SpuId:       spu.Id,
		AttrValues:  attrValues,
		LogoUrl:     logoUrl,
		MainUrl:     mainUrl,
		BannerUrl:   BannelUrl,
		SellPrice:   sellPrice,
		MarketPrice: marketPrice,
		Stock:       stock,
	}
	if err = models.Db.Create(model).Error; err != nil {
		return
	}
	skuResp.composeField(&sku, model)
	return
}

func (skuResp *SkuResp) Update(sku entities.Sku) (err error) {
	model := &models.Sku{
		Name:        sku.Name,
		Desc:        sku.Desc,
		AttrValues:  sku.AttrValues,
		LogoUrl:     sku.LogoUrl,
		MainUrl:     sku.MainUrl,
		BannerUrl:   sku.BannerUrl,
		SellPrice:   sku.SellPrice,
		MarketPrice: sku.MarketPrice,
		Stock:       sku.Stock,
	}
	err = models.Db.Model(model).Where("id =?", sku.Id).Updates(model).Error
	return
}

func (skuResp *SkuResp) Get(id uint) (sku entities.Sku, err error) {
	model := new(models.Sku)
	if err = models.Db.Where("id =?", id).First(model).Error; err != nil {
		return
	}
	spuResp := new(SpuResp)
	spu, err := spuResp.Get(model.SpuId)
	if err != nil {
		return
	}
	sku.Spu = spu
	skuResp.composeField(&sku, model)
	return
}

func (skuResp *SkuResp) GetSkuByAttrValues(spuId uint, attrValues string) (sku entities.Sku, err error) {
	var skus []models.Sku
	if err = models.Db.Where("spu_id = ?", spuId).Find(&skus).Error; err != nil {
		return
	}
	var searchAttrValues []map[string]uint
	if err = json.Unmarshal([]byte(attrValues), &searchAttrValues); err != nil {
		return
	}
	var searchValueIds []uint
	for _, rearchAttrValue := range searchAttrValues {
		searchValueIds = append(searchValueIds, rearchAttrValue["value_id"])
	}
	sort.Slice(searchValueIds, func(i, j int) bool {
		return i < j
	})
	var skuAttrValues []map[string]uint
	var isFind = false
	for _, skuModel := range skus {
		json.Unmarshal([]byte(skuModel.AttrValues), &skuAttrValues)
		var skuValueIds []uint
		for _, skuAttrValue := range skuAttrValues {
			skuValueIds = append(skuValueIds, skuAttrValue["value_id"])
		}
		sort.Slice(skuValueIds, func(i, j int) bool {
			return i < j
		})
		if util.CompareSliceUint(searchValueIds, skuValueIds) {
			skuResp.composeField(&sku, &skuModel)
			isFind = true
			break
		}
	}
	if !isFind {
		err = errors.New("没有找到对应的SKU商品")
	}
	return
}

func (skuResp *SkuResp) composeField(sku *entities.Sku, model *models.Sku) {
	sku.Id = model.ID
	sku.Name = model.Name
	sku.Desc = model.Desc
	sku.LogoUrl = model.LogoUrl
	sku.MainUrl = model.MainUrl
	sku.BannerUrl = model.BannerUrl
	sku.SellPrice = model.SellPrice
	sku.MarketPrice = model.MarketPrice
	sku.Stock = model.Stock
	sku.Status = model.Status
	sku.CreatedAt = model.CreatedAt
}

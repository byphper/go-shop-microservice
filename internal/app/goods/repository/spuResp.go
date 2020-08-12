package repository

import (
	"github.com/pkg/errors"
	"go-shop-microservice/internal/app/goods/entities"
	"go-shop-microservice/internal/app/goods/models"
	"strconv"
	"strings"
)

type SpuResp struct {
}

func (spuResp *SpuResp) Create(name, desc, logoUrl, mainUrl, BannelUrl, unit, attrIds string, cid, bid, sellPrice, marketPrice uint) (spu entities.Spu, err error) {
	cateResp := new(CategoryResp)
	cate, err := cateResp.Get(cid)
	if err != nil {
		return spu, errors.Wrap(err, "invalid category")
	}
	spu.Category = cate
	brandResp := new(BrandResp)
	brand, err := brandResp.Get(bid)
	if err != nil {
		return spu, errors.Wrap(err, "invalid brand")
	}
	spu.Brand = brand
	attrIdsSlice := strings.Split(attrIds, "-")
	if len(attrIdsSlice) == 0 {
		return spu, errors.Wrap(err, "invalid attr")
	}
	model := &models.Spu{
		Name:        name,
		BrandId:     bid,
		CategoryId:  cid,
		Desc:        desc,
		AttrIds:     attrIds,
		LogoUrl:     logoUrl,
		MainUrl:     mainUrl,
		BannerUrl:   BannelUrl,
		SellPrice:   sellPrice,
		MarketPrice: marketPrice,
		Unit:        unit,
	}
	if err = models.Db.Create(model).Error; err != nil {
		return
	}
	spuResp.composeField(&spu, model)
	return
}

func (spuResp *SpuResp) Update(spu entities.Spu) (err error) {
	attrIdsSlice := strings.Split(spu.AttrIds, "-")
	if len(attrIdsSlice) == 0 {
		return errors.Wrap(err, "invalid attr")
	}
	model := &models.Spu{
		Name:        spu.Name,
		BrandId:     spu.Brand.Id,
		CategoryId:  spu.Category.Id,
		Desc:        spu.Desc,
		AttrIds:     spu.AttrIds,
		LogoUrl:     spu.LogoUrl,
		MainUrl:     spu.MainUrl,
		BannerUrl:   spu.BannerUrl,
		SellPrice:   spu.SellPrice,
		MarketPrice: spu.MarketPrice,
		Unit:        spu.Unit,
	}
	err = models.Db.Model(model).Where("id =?", spu.Id).Updates(model).Error
	return
}

func (spuResp *SpuResp) Get(id uint) (spu entities.Spu, err error) {
	model := new(models.Spu)
	if err = models.Db.Where("id =?", id).First(model).Error; err != nil {
		return
	}
	cateResp := new(CategoryResp)
	cate, _ := cateResp.Get(model.CategoryId)
	spu.Category = cate
	brandResp := new(BrandResp)
	brand, _ := brandResp.Get(model.BrandId)
	spu.Brand = brand
	attrResp := new(AttributeResp)
	attrIdsSlice := strings.Split(model.AttrIds, "-")
	var attrs []entities.Attribute
	for _, v := range attrIdsSlice {
		int64Tmp, _ := strconv.ParseUint(v, 10, 32)
		attr, _ := attrResp.Get(uint(int64Tmp))
		attrs = append(attrs, attr)
	}
	spu.Attrs = attrs
	spuResp.composeField(&spu, model)
	return
}

func (spuResp *SpuResp) composeField(spu *entities.Spu, model *models.Spu) {
	spu.Id = model.ID
	spu.Name = model.Name
	spu.Desc = model.Desc
	spu.LogoUrl = model.LogoUrl
	spu.MainUrl = model.MainUrl
	spu.BannerUrl = model.BannerUrl
	spu.SellPrice = model.SellPrice
	spu.MarketPrice = model.MarketPrice
	spu.Unit = model.Unit
	spu.Status = model.Status
	spu.CreatedAt = model.CreatedAt
}

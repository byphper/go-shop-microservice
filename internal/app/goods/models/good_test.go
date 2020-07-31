package models

import (
	"testing"
)

var gd = &GoodDaoImpl{}

func TestGoodDaoImpl_CreateAttr(t *testing.T) {
	attr, err := gd.CreateAttr("尺码", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(attr)
}

func TestGoodDaoImpl_CreateAttrValue(t *testing.T) {
	value, err := gd.CreateAttrValue(2, "36", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(value)
}

func TestGoodDaoImpl_CreateBrand(t *testing.T) {
	value, err := gd.CreateBrand("Lining", "", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(value)
}

func TestGoodDaoImpl_CreateSpu(t *testing.T) {
	spuModel := GoodSpu{
		Name:           "李宁运动鞋",
		GoodBrandId:    2,
		GoodCategoryId: 1,
		Desc:           "好鞋",
		LogoUrl:        "",
		MainUrl:        "",
		BannerUrl:      "",
		SellPrice:      1000,
		MarketPrice:    1500,
		sales:          0,
		Unit:           "双",
		Status:         0,
	}
	spu, err := gd.CreateSpu(&spuModel)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(spu)
}

func TestGoodDaoImpl_GetSpuById(t *testing.T) {
	spu, err := gd.GetSpuById(3)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(spu.Name, spu.GoodBrand.Name, spu.GoodCategory.Name)
}

func TestGoodDaoImpl_CreateSku(t *testing.T) {
	sku := &GoodSku{
		GoodSpuId:       3,
		Name:            "李宁运动鞋骚红款",
		Desc:            "不臭脚",
		LogoUrl:         "",
		MainUrl:         "",
		BannerUrl:       "",
		GoodsAttrValues: "1,3",
		SellPrice:       5000,
		MarketPrice:     5500,
		Quantity:        9999,
		Status:          0,
	}
	sku, err := gd.CreateSku(sku)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sku)
}

func TestGoodDaoImpl_UpdateSpu(t *testing.T) {
	spu, err := gd.GetSpuById(3)
	if err != nil {
		t.Fatal(err)
	}
	spu.SellPrice = 2000
	spu.Desc = "真是一双好鞋啊"
	spu.Status = StatusEnable
	_, err = gd.UpdateSpu(spu)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("修改成功")
}

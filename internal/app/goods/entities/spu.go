package entities

import "time"

type Spu struct {
	Id          uint
	Name        string
	Brand       Brand
	Category    Category
	AttrIds     string
	Attrs        []Attribute
	Desc        string
	LogoUrl     string
	MainUrl     string
	BannerUrl   string
	SellPrice   uint
	MarketPrice uint
	sales       uint
	Unit        string
	Status      uint8
	CreatedAt   time.Time
}

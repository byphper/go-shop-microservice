package entities

import "time"

type Sku struct {
	Id uint
	Spu         Spu
	Name        string
	Desc        string
	LogoUrl     string
	MainUrl     string
	BannerUrl   string
	SellPrice   uint
	MarketPrice uint
	Stock       uint
	Status      uint8
	CreatedAt time.Time
}

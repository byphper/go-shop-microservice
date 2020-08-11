package main

import "go-shop-microservice/internal/app/goods/models"

func main() {
	models.Db.AutoMigrate(&models.Sku{})

}

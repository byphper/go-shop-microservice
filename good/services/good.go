package services

import (
	"go-shop-microservice/good/models"
)

type GoodService struct {
	userDao *models.GoodDaoImpl
}

func (gs GoodService) GetSpuDetail(id int) (*models.GoodSpu, error) {
	spu, err := gs.userDao.GetSpuById(uint(id))
	if err != nil {
		return nil, err
	}

}

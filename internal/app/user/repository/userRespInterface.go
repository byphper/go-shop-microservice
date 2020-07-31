package repository

import (
	"go-shop-microservice/internal/app/user/entities"
)

type UserRespInterface interface {
	Get(id int) (entities.User, error)
	GetByEmail(email string) (entities.User, error)
	IsExist(email string) bool
	Create(email string, name string, avatar string, pwd string) (entities.User, error)
	Update(user entities.User) error
}

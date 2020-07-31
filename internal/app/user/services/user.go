package services

import (
	"go-shop-microservice/internal/app/user/entities"
	"go-shop-microservice/internal/app/user/repository"
)

//用户领域服务
type User struct {
	userResp repository.UserRespInterface
}

//通过ID获取用户信息
func (us *User) GetUser(id int) (entities.User, error) {
	return us.userResp.Get(id)
}

//通过email获取用户信息
func (us *User) GetUserByEmail(email string) (entities.User, error) {
	return us.userResp.GetByEmail(email)
}

//创建用户
func (us *User) CreateUser(email string, name string, avatar string, pwd string) (entities.User, error) {
	return us.userResp.Create(email, name, avatar, pwd)
}

//更新用户
func (us *User) UpdateUser(user entities.User) error {
	return us.userResp.Update(user)
}

package repository

import (
	"go-shop-microservice/internal/app/user/entities"
	"go-shop-microservice/internal/app/user/models"
)

//用户仓储
type UserResp struct {
}

func (u *UserResp) IsExist(email string) bool {
	userModel := &models.User{}
	if err := models.Db.Select("id").Where("email =?", email).First(userModel).Error; err != nil {
		return false
	}
	return true
}

func (u *UserResp) Get(id int) (userEntity entities.User, err error) {
	userModel := &models.User{}
	if err = models.Db.Where("id =?", id).First(userModel).Error; err != nil {
		return
	}
	u.fillUserField(&userEntity, userModel)
	return
}

func (u *UserResp) GetByEmail(email string) (userEntity entities.User, err error) {
	userModel := &models.User{}
	if err = models.Db.Where("email =?", email).First(userModel).Error; err != nil {
		return
	}
	u.fillUserField(&userEntity, userModel)
	return
}

func (u *UserResp) fillUserField(user *entities.User, model *models.User) {
	user.Id = model.ID
	user.Name = model.Name
	user.Email = model.Email
	user.Avatar = model.Avatar
	user.CreatedAt = model.CreatedAt
	user.VerifiedAt = model.VerifiedAt
	user.Pwd = model.Pwd
}

func (u *UserResp) Create(email string, name string, avatar string, pwd string) (userEntity entities.User, err error) {
	userModel := &models.User{
		Name:   name,
		Email:  email,
		Pwd:    pwd,
		Avatar: avatar,
	}
	if err = models.Db.Create(userModel).Error; err != nil {
		return
	}
	userEntity.Id = userModel.ID
	userEntity.Name = userModel.Name
	userEntity.Email = userModel.Email
	userEntity.Avatar = userModel.Avatar
	userEntity.CreatedAt = userModel.CreatedAt
	return
}

func (u *UserResp) Update(user entities.User) (err error) {
	userModel := &models.User{}
	err = models.Db.Where("id =?", user.Id).First(userModel).Error
	return
}

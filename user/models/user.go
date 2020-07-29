package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name       string `gorm:"type:varchar(10);unique_index"`
	Email      string `gorm:"type:varchar(50);index:ik_email"`
	Pwd        string `gorm:"type:varchar(64)"`
	Avatar     string `gorm:"varchar(255)"`
	VerifiedAt *time.Time
}

type UserDaoImpl struct {
}

func (u *UserDaoImpl) Get(id int, fields ...string) (*User, error) {
	um := &User{}
	if fields == nil {
		fields = []string{"*"}
	}
	if err := Db.Select(fields).Where("id =?", id).First(um).Error; err != nil {
		return nil, err
	}
	return um, nil
}

func (u *UserDaoImpl) GetByEmail(email string, fields ...string) (*User, error) {
	um := &User{}
	if fields == nil {
		fields = []string{"*"}
	}
	if err := Db.Select(fields).Where("email =?", email).First(um).Error; err != nil {
		return nil, err
	}
	return um, nil
}

func (u *UserDaoImpl) Create(user *User) (*User, error) {
	if err := Db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserDaoImpl) Update(user *User) error {
	return Db.Model(user).Update(user).Error
}

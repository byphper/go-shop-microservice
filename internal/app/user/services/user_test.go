package services

import (
	"go-shop-microservice/internal/app/user/repository"
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {
	userService := User{userResp: &repository.UserResp{}}
	email, name, pwd, avatar := "b@g.com", "bob", "123123", "default"
	user, err := userService.CreateUser(email, name, avatar, pwd)
	if err != nil {
		t.Log(err)
	}
	t.Log(user)
}

func TestUserService_GetUser(t *testing.T) {
	userService := User{userResp: &repository.UserResp{}}
	user, err := userService.GetUser(1)
	if err != nil {
		t.Log(err)
	}
	t.Log(user)
}

func TestUserService_GetUserByEmail(t *testing.T) {
	userService := User{userResp: &repository.UserResp{}}
	user, err := userService.GetUserByEmail("b@g.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}

func TestUserService_UpdateUser(t *testing.T) {
	userService := User{userResp: &repository.UserResp{}}
	user, err := userService.GetUser(1)
	if err != nil {
		t.Fatal(err)
	}
	user.Avatar = "http://www.taobao.com"
	err = userService.UpdateUser(user)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("更新成功")
}

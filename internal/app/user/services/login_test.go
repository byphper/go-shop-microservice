package services

import (
	"go-shop-microservice/internal/app/user/repository"
	"testing"
)

func TestLoginService_RegisterByEmail(t *testing.T) {
	loginService := Login{userResp: &repository.UserResp{}}
	user, err := loginService.RegisterByEmail("b@qq.com", "123123", "andy", "default")
	if err != nil {
		t.Error(err)
	}
	t.Log(user)
}

func TestLoginService_LoginByEmail(t *testing.T) {
	loginService := Login{userResp: &repository.UserResp{}}
	email, pwd := "b@qq.com", "123123"
	user, err := loginService.LoginByEmail(email, pwd)
	if err != nil {
		t.Error(err)
	}
	t.Log(user)
}

func TestLoginService_SendRegisterEmail(t *testing.T) {
	loginService := Login{userResp: &repository.UserResp{}}
	email := "b@qq.com"
	err := loginService.SendRegisterEmail(email)
	if err != nil {
		t.Error(err)
	}
	t.Log("发送成功")
}

func TestLoginService_VerifyEmail(t *testing.T) {
	loginService := Login{userResp: &repository.UserResp{}}
	email := "b@qq.com"
	err := loginService.VerifyEmail(email,"1596174636","73c0cef9a4b3ac646d708c828c78be324334af08")
	if err != nil {
		t.Error(err)
	}
	t.Log("验证成功")
}

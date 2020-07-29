package services

import (
	"testing"
	usermodels "xxswkl/user/models"
	vo "xxswkl/user/vo"
)

func TestUserService_CreateUser(t *testing.T) {
	pwd, pwdErr := GeneratePwd("123456")
	if pwdErr != nil {
		t.Fatal(pwdErr)
	}
	ud := &usermodels.UserDaoImpl{}
	us := UserService{userDao: ud}
	user := vo.User{
		Name:   "kaki",
		Email:  "kaki@gmail.com",
		Avatar: "http://cdn.baidu.com/andy.jgp",
		Pwd:    string(pwd),
	}
	model, err := us.CreateUser(user)
	if err != nil {
		t.Log(err)
	}
	t.Log(model)
}

func TestUserService_GetUser(t *testing.T) {
	ud := &usermodels.UserDaoImpl{}
	us := UserService{userDao: ud}
	user, err := us.GetUser(5)
	if err != nil {
		t.Log(err)
	}
	t.Log(user)
}

func TestUserService_GetUserByEmail(t *testing.T) {
	ud := &usermodels.UserDaoImpl{}
	us := UserService{userDao: ud}
	user, err := us.GetUserByEmail("kenji@gmail.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}

func TestUserService_UpdateUser(t *testing.T) {
	ud := &usermodels.UserDaoImpl{}
	us := UserService{userDao: ud}
	user, err := us.GetUser(5)
	if err != nil {
		t.Fatal(err)
	}
	user.Avatar = "http://www.taobao.com"
	err = us.UpdateUser(user)
	if err != nil {
		t.Fatal("更新失败")
	}
}

func TestUserService_LoginByEmail(t *testing.T) {
	ud := &usermodels.UserDaoImpl{}
	us := UserService{userDao: ud}
	user, err := us.LoginByEmail("andy@gmail.com", "123456")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}

func TestUserService_RegisterByEmail(t *testing.T) {
	ud := &usermodels.UserDaoImpl{}
	us := UserService{userDao: ud}
	user, err := us.RegisterByEmail("kenji@gmail.com", "123456", "Kenji", "http://www.sohu.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}

func TestUserService_SendRegisterEmail(t *testing.T) {
	us := UserService{}
	us.SendRegisterEmail("andy@gmail.com")
}

func TestUserService_VerifyEmail(t *testing.T) {
	ud := &usermodels.UserDaoImpl{}
	us := UserService{userDao: ud}
	err := us.VerifyEmail("andy@gmail.com", "1589027681", "32717ea0a591c287f7d37ca92580d4853ab42134")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("验证成功")
}

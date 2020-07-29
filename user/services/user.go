package services

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"go-shop-microservice/user/models"
	"strconv"
	"time"
)

const EmailActiveToken = "baoyong"

//用户服务
type UserService struct {
	userDao *models.UserDaoImpl
}

//通过ID获取用户信息
func (us *UserService) GetUser(id int, fields ...string) (*models.User, error) {
	return us.userDao.Get(id, fields...)
}

//通过email获取用户信息
func (us *UserService) GetUserByEmail(email string, fields ...string) (*models.User, error) {
	return us.userDao.GetByEmail(email, fields...)
}

//创建用户
func (us *UserService) CreateUser(user *models.User) (*models.User, error) {
	return us.userDao.Create(user)
}

//更新用户
func (us *UserService) UpdateUser(user *models.User) error {
	return us.userDao.Update(user)
}

//email登录
func (us *UserService) LoginByEmail(email string, pwd string) (*models.User, error) {
	user, err := us.userDao.GetByEmail(email)
	if err != nil {
		return user, errors.New("无效的账号")
	}
	ok, err := ValidatePwd(pwd, user.Pwd)
	if !ok {
		return user, errors.New("密码错误")
	}
	return user, nil
}

//emial注册用户
func (us *UserService) RegisterByEmail(email string, pwd string, name string, avatar string) (*models.User, error) {
	user := &models.User{
		Name:   name,
		Email:  email,
		Avatar: avatar,
	}
	hashedPwd, pwdErr := GeneratePwd(pwd)
	if pwdErr != nil {
		return nil, errors.New("系统错误")
	}
	user.Pwd = string(hashedPwd)
	registeredUser, err := us.CreateUser(user)
	if err != nil {
		return nil, errors.New("注册失败")
	}
	return registeredUser, nil
}

//发送注册邮件
func (us *UserService) SendRegisterEmail(email string) error {
	now := time.Now().Unix()
	tokenStr := strconv.FormatInt(int64(now), 10) + email + EmailActiveToken
	sha := sha1.New()
	sha.Write([]byte(tokenStr))
	token := hex.EncodeToString(sha.Sum(nil))
	url := fmt.Sprintf("user/email/verify?email=%v&timestamp=%v&token=%v", email, now, token)
	fmt.Println(url)
	return nil
}

//验证注册邮件
func (us *UserService) VerifyEmail(email string, timeParam string, tokenParam string) error {
	now := time.Now().Unix()
	timeInt, err := strconv.ParseInt(timeParam, 10, 64)
	if err != nil {
		return errors.New("无效的参数")
	}
	if now > timeInt+1800 {
		return errors.New("验证链接已过期")
	}
	tokenStr := timeParam + email + EmailActiveToken
	sha := sha1.New()
	sha.Write([]byte(tokenStr))
	token := hex.EncodeToString(sha.Sum(nil))
	if token != tokenParam {
		return errors.New("无效的签名")
	}
	user, err := us.GetUserByEmail(email, "id", "verified_at")
	if err != nil {
		return errors.New("无效的用户邮箱")
	}
	if !user.VerifiedAt.IsZero() {
		return nil
	}
	verifiedTime := time.Now()
	user.VerifiedAt = &verifiedTime
	err = us.UpdateUser(user)
	if err != nil {
		return errors.New("验证失败")
	}
	return nil
}

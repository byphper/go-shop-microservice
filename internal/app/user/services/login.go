package services

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"go-shop-microservice/internal/app/user/entities"
	"go-shop-microservice/internal/app/user/repository"
	"go-shop-microservice/internal/pkg/log"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

const EmailActiveToken = "baoyong"

//登录注册服务
type Login struct {
	UserResp repository.UserRespInterface
}

//email登录
func (loginService *Login) LoginByEmail(email string, pwd string) (entities.User, error) {
	user, err := loginService.UserResp.GetByEmail(email)
	if err != nil {
		return user, errors.New("无效的账号")
	}
	ok, err := loginService.validatePwd(pwd, user.Pwd)
	if !ok {
		return user, errors.New("密码错误")
	}
	return user, nil
}

//emial注册用户
func (loginService *Login) RegisterByEmail(email string, pwd string, name string, avatar string) (entities.User, error) {
	user := entities.User{}
	hashedPwd, err := loginService.generatePwd(pwd)
	if err != nil {
		log.Logger.Error(err.Error())
		return user, errors.New("系统错误")
	}
	ok := loginService.UserResp.IsExist(email)
	if ok {
		return user, errors.New("邮箱已经存在")
	}
	user, err = loginService.UserResp.Create(email, name, avatar, string(hashedPwd))
	if err != nil {
		return user, errors.New("注册失败")
	}
	return user, nil
}

//发送注册邮件
func (loginService *Login) SendRegisterEmail(email string) error {
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
func (loginService *Login) VerifyEmail(email string, timeParam string, tokenParam string) error {
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
	user, err := loginService.UserResp.GetByEmail(email)
	if err != nil {
		return errors.New("无效的用户邮箱")
	}
	if nil == user.VerifiedAt {
		return nil
	}
	verifiedTime := time.Now()
	user.VerifiedAt = &verifiedTime
	err = loginService.UserResp.Update(user)
	if err != nil {
		return errors.New("验证失败")
	}
	return nil
}

func (loginService *Login) generatePwd(sourcePwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(sourcePwd), bcrypt.DefaultCost)
}

func (loginService *Login) validatePwd(sourcePwd string, hashedPwd string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(sourcePwd)); err != nil {
		return false, fmt.Errorf("校验失败%w", err)
	}
	return true, nil
}

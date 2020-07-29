package services

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePwd(sourcePwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(sourcePwd), bcrypt.DefaultCost)
}

func ValidatePwd(sourcePwd string, hashedPwd string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(sourcePwd)); err != nil {
		return false, fmt.Errorf("校验失败%w", err)
	}
	return true, nil
}

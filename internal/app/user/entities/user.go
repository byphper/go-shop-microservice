package entities

import "time"

//用户实体
type User struct {
	Id         uint
	Email      string
	Name       string
	Avatar     string
	Pwd        string
	CreatedAt  time.Time
	VerifiedAt *time.Time
}


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


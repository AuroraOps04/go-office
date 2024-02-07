package model

import (
	"github.com/AuroraOps04/go-office/model/field"
)

type User struct {
	field.BaseModel
	Username   string `json:"username"`
	Password   string `json:"password"   gorm:"type:char(60)"`
	NickName   string `json:"nickname"`
	Phone      string `json:"phone"`
	Department string `json:"department"`
	Status     int8   `json:"status"`
	Role       int8   `json:"role"`
	OpenID     string `json:"openid"`
}

func (u *User) TableName() string {
	return "user"
}

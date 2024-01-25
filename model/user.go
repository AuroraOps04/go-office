package model

import (
	"github.com/AuroraOps04/go-office/model/field"
)

type User struct {
	field.BaseModel
	Username   string `json:"username"`
	Password   string `json:"password"`
	NickName   string `json:"nickname"`
	Phone      string
	Department string
	Status     int8
	Role       int8
	OpenID     string
}

func (u *User) TableName() string {
	return "user"
}

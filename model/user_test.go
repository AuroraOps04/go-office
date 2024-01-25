package model_test

import (
	"testing"

	"github.com/AuroraOps04/go-office/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestUser(t *testing.T) {
	db, err := gorm.Open(mysql.Open("mysql"))
	if err != nil {
		t.Fatal(err)
	}
	_ = db.AutoMigrate(&model.User{})
}

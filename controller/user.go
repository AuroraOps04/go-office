package controller

import (
	"errors"
	"fmt"

	"github.com/AuroraOps04/go-office/controller/base"
	"github.com/AuroraOps04/go-office/global"
	"github.com/AuroraOps04/go-office/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userController struct {
	*base.BaseController
}

func NewUser() *userController {
	return &userController{}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *userController) CurrentUser(ctx *gin.Context) {
	session := sessions.Default(ctx)
	var currentUser model.User
	uid, ok := session.Get("currentUser").(uint)
	if !ok {
		c.ErrorMsg(ctx, "没有登陆请登录")
		return
	}
	global.Db.Model(&model.User{}).Take(&currentUser, uid)
	c.SuccessData(ctx, currentUser)
}

func (c *userController) Login(ctx *gin.Context) {
	var r LoginRequest
	if err := ctx.Copy().ShouldBindJSON(&r); err != nil {
		c.ErrorMsg(ctx, "请输入用户名或者密码")
		return
	}
	db := global.Db
	var u model.User
	if tx := db.Model(&model.User{}).First(&u, "username = ?", r.Username); tx.Error != nil {
		c.ErrorMsg(ctx, "账号或密码错误")
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(r.Password)) != nil {
		c.ErrorMsg(ctx, "账号或密码错误")
		return
	}
	// login success
	session := sessions.Default(ctx)
	session.Set("userId", u.ID)
	// INFO: Don't forget session.Save()
	err := session.Save()
	if err != nil {
		fmt.Printf("save session error:%s", err.Error())
	}
	// save to session
	c.SuccessData(ctx, u)
}

func (c *userController) Register(ctx *gin.Context) {
	var u model.User
	if err := ctx.ShouldBindJSON(&u); err != nil {
		c.ErrorMsg(ctx, "参数错误")
		return
	}
	// 判断用户名是否存在
	var exituser model.User
	if tx := global.Db.Model(&model.User{}).First(&exituser, "username = ?", u.Username); !errors.Is(
		tx.Error,
		gorm.ErrRecordNotFound,
	) {
		c.ErrorMsg(ctx, "用户已存在")
		return
	}
	// 密码加密
	if hashedPwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost); err != nil {
		c.ErrorMsg(ctx, "加密密码失败")
		return
	} else {
		u.Password = string(hashedPwd)
	}

	// 保存用户
	if tx := global.Db.Model(&model.User{}).Create(&u); tx.Error != nil || u.ID == 0 {
		c.ErrorMsg(ctx, "创建用户失败")
		return
	}
	c.Success(ctx)
}

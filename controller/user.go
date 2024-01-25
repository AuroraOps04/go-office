package controller

import (
	"github.com/AuroraOps04/go-office/controller/base"
	"github.com/gin-gonic/gin"
)

type userController struct {
	*base.BaseController
}

func (c *userController) CurrentUser(ctx *gin.Context) {
}

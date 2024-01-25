package base

import "github.com/gin-gonic/gin"

type BaseController struct{}

func (c *BaseController) SuccessMsg(ctx *gin.Context, msg string) {
}

func (c *BaseController) SuccessData(ctx *gin.Context, data string) {
}

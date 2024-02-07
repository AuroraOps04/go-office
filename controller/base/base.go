package base

import (
	"github.com/AuroraOps04/go-office/model/response"
	"github.com/gin-gonic/gin"
)

type BaseController struct{}


func (c *BaseController) ErrorMsg(ctx *gin.Context, msg string) {
  ctx.JSON(200, response.Response{
  	Success:      false,
  	Data:         nil,
  	Total:        0,
  	ErrorMessage: msg,
  	ErrorCode:    400,
  	ShowType:     response.ERROR,
  })
}

func (c *BaseController) SuccessData(ctx *gin.Context, data interface{}) {
  ctx.JSON(200, response.Response{
  	Success:      true,
  	Data:         data,
  	Total:        0,
  	ErrorMessage: "",
  	ErrorCode:    0,
  	ShowType:     response.SILENT,
  })
}

func (c *BaseController) SuccessMsg(ctx *gin.Context, msg string) {
  ctx.JSON(200, response.Response{
  	Success:      true,
  	Data:         msg,
  	Total:        0,
  	ErrorMessage: "",
  	ErrorCode:    0,
  	ShowType:     response.SILENT,
  })
}
func (c *BaseController) Success(ctx *gin.Context) {
  ctx.JSON(200, response.Response{
  	Success:      true,
  	Data:         "ok",
  	Total:        0,
  	ErrorMessage: "",
  	ErrorCode:    0,
  	ShowType:     response.SILENT,
  })
}

func (c *BaseController) SuccessPage(ctx *gin.Context, data interface{}, total int64) {
  ctx.JSON(200, response.Response{
  	Success:      true,
  	Data:         data,
  	Total:        total,
  	ErrorMessage: "",
  	ErrorCode:    0,
  	ShowType:     response.SILENT,
  })
}

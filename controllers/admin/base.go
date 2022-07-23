package admin

import "github.com/gin-gonic/gin"

type BaseController struct {
}

func (con BaseController) Success(ctx *gin.Context) {
	ctx.String(200, "2312312")
}

func (con BaseController) Error(ctx *gin.Context) {
	ctx.String(200, "2312312")
}

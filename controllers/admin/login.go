package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
import "xiaomi-service/models"

type LoginController struct {
}
type login struct {
	CaptchaId   string `json:"captchaId" binding:"required"`
	CaptchaCode string `json:"captchaCode" binding:"required"`
}

func (con LoginController) GetCaptcha(ctx *gin.Context) {
	id, bs64, err := models.MakeCaptcha()
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"captchaId":  id,
		"captchaImg": bs64,
	})
}

func (con LoginController) Login(ctx *gin.Context) {

	data := login{}
	err := ctx.BindJSON(&data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		fmt.Println(data)
		flag := models.VerifyCaptcha(data.CaptchaId, data.CaptchaCode)
		if flag {
			ctx.String(http.StatusOK, "验证成功")
		} else {
			ctx.String(http.StatusOK, "验证失败")
		}
	}

}

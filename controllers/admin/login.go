package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
import "xiaomi-service/models"

type LoginController struct {
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
	type login struct {
		captchaId  string
		verifyCode string
	}

	json := login{}
	err := ctx.BindJSON(&json)
	log.Printf("%v", json)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		flag := models.VerifyCaptcha(json.captchaId, json.verifyCode)
		if flag {
			ctx.String(http.StatusOK, "验证成功")
		} else {
			ctx.String(http.StatusOK, "验证失败")
		}
	}

	//captchaId := ctx.PostForm("captchaId")
	//verifyCode := ctx.PostForm("verifyCode")
	//fmt.Println("3213")
	//fmt.Println(captchaId)
	//fmt.Println(verifyCode)

}

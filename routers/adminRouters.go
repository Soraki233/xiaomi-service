package routers

import (
	"github.com/gin-gonic/gin"
	"xiaomi-service/controllers/admin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/getCaptcha", admin.LoginController{}.GetCaptcha)
		adminRouters.POST("/login", admin.LoginController{}.Login)
	}
}

package routers

import (
	"github.com/gin-gonic/gin"
	"xiaomi-service/controllers/admin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/rest", admin.BaseController{}.Success)
	}
}

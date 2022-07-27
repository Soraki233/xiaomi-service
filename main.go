package main

import (
	"github.com/gin-gonic/gin"
	"xiaomi-service/routers"
)

func main() {
	r := gin.Default()
	//r.Use(utils.Core())
	routers.AdminRoutersInit(r)
	r.Run(":8080")
}

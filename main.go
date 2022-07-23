package main

import (
	"github.com/gin-gonic/gin"
	"xiaomi-service/routers"
)

func main() {
	r := gin.Default()
	routers.AdminRoutersInit(r)
	r.Run(":8080")
}

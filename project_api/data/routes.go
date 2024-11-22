package data

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine) {
	group := router.Group("/data")

	group.GET("", ShowAllData)
	group.GET("/filter", FilterAndOrderData)
}

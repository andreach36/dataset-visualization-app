package data

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine){
	group := router.Group("/data")

	group.GET("", ShowAllData)
	group.GET("/filter/education", FilterDataByEducation)
	group.GET("/filter/marital_status", FilterDataByMaritalStatus)
	group.GET("/filter/occupation", FilterDataByOcupation)
	group.GET("/filter/age", FilterDataByAge)
	group.GET("/filter/income", FilterDataByIncome)
	group.GET("/order/asc", OrderDataByAscendant)
	group.GET("/order/desc", OrderDataByDescendant)

}
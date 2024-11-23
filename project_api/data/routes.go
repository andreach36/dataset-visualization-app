package data

import (
	"trucode3-challenge-final/project_api/shared"

	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine) {
	group := router.Group("/data")
	group.Use(shared.AuthenticateSession())

	group.GET("", ShowAllData)

	group.PUT("/filters", UpdateUserFilter)  // Actualiza un filtro
	group.GET("/filters", GetUserFilter)     // Obtener los filtros del usuario
	group.GET("/filter", FilterAndOrderData) // Filtra la data
}

package data

import (
	"trucode3-challenge-final/project_api/shared"

	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine) {
	group := router.Group("/data")
	group.Use(shared.AuthenticateSession())

	group.GET("/filters", InitializeUserFilters) // Inicializa al usuario con filtros predeterminados

	group.GET("", ShowAllData)

	group.PUT("/filters/:id", UpdateUserFilter) // Actualiza un filtro
	group.GET("/filter", FilterAndOrderData)    // Filtra la data
}

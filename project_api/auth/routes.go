package auth

import (
	"project_api/shared"

	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine) {
	group := router.Group("/auth")
	group.POST("/register", Register)
	group.POST("/login", Login)
	group.DELETE("/logout", shared.AuthenticateSession(), Logout)
}

package infrastructure

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/api/auth"
)

func SetupRoutes(userHandler auth.Handler) *gin.Engine {
	router := gin.Default()

	a := router.Group("/auth")
	a.POST("/register", userHandler.RegisterUserHandler)
	a.POST("/login", userHandler.LoginUserHandler)

	return router
}

package app

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/domains/user"
)

func SetupRoutes(userHandler user.UserHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/user", userHandler.GetAllUserHandler)
	router.GET("/user/:id", userHandler.GetUserByIDHandler)
	router.POST("/user", userHandler.CreateUserHandler)
	router.PUT("/user/:id", userHandler.UpdateUserHandler)
	router.DELETE("/user/:id", userHandler.DeleteUserHandler)

	return router
}

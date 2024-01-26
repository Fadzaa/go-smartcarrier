package infrastructure

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/api/user"
)

func SetupRoutes(userHandler user.UserHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/user", userHandler.GetAllUserHandler)
	router.GET("/user/:id", userHandler.GetUserByIDHandler)
	router.POST("/user", userHandler.CreateUserHandler)
	router.PUT("/user/:id", userHandler.UpdateUserHandler)
	router.DELETE("/user/:id", userHandler.DeleteUserHandler)

	//router.GET("/job", jobHandler.GetAllJobHandler)
	//router.GET("/job/:id", jobHandler.GetJobByIDHandler)
	//router.POST("/job", jobHandler.CreateJobHandler)
	//router.PUT("/job/:id", jobHandler.UpdateJobHandler)
	//router.DELETE("/job/:id", jobHandler.DeleteJobHandler)

	return router
}

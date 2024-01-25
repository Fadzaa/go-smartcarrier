package main

import (
	"github.com/gin-gonic/gin"
	userHandler "go-gin-api/app/handlers/user"
	"go-gin-api/domains/user"
	"go-gin-api/infrastructure"
)

func main() {
	//Open Database Connection
	db := infrastructure.ConnectToDatabase()

	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	handlerUser := userHandler.NewHandlerUser(userService)

	//Initialize Gin Router
	router := gin.Default()

	u := router.Group("/user")
	u.GET("/", handlerUser.GetAllUserHandler)
	u.GET("/:id", handlerUser.GetUserByIDHandler)
	u.POST("/", handlerUser.CreateUserHandler)
	u.PUT("/:id", handlerUser.UpdateUserHandler)
	u.DELETE("/:id", handlerUser.DeleteUserHandler)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}

}

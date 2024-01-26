//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go-gin-api/app"
	"go-gin-api/domains/user"
)

var userSet = wire.NewSet(
	user.NewUserRepository,
	wire.Bind(new(user.UserRepository), new(*user.UserRepositoryImpl)),
	user.NewUserService,
	wire.Bind(new(user.UserService), new(*user.UserServiceImpl)),
	user.NewHandlerUser,
	wire.Bind(new(user.UserHandler), new(*user.UserHandlerImpl)),
)

func InitializedUser() *gin.Engine {
	wire.Build(
		app.ConnectToDatabase,
		userSet,
		app.SetupRoutes,
	)
	return nil
}

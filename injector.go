//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	user2 "go-gin-api/api/user"
	"go-gin-api/domain/user"
	"go-gin-api/infrastructure"
)

var userSet = wire.NewSet(
	user.NewUserRepository,
	wire.Bind(new(user.UserRepository), new(*user.UserRepositoryImpl)),
	user.NewUserService,
	wire.Bind(new(user.UserService), new(*user.UserServiceImpl)),
	user2.NewHandlerUser,
	wire.Bind(new(user2.UserHandler), new(*user2.UserHandlerImpl)),
)

func InitializedApp() *gin.Engine {
	wire.Build(
		infrastructure.ConnectToDatabase,
		userSet,
		infrastructure.SetupRoutes,
	)
	return nil
}

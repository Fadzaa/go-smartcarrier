// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go-gin-api/app"
	"go-gin-api/domains/user"
)

// Injectors from injector.go:

func InitializedUser() *gin.Engine {
	db := app.ConnectToDatabase()
	userRepositoryImpl := user.NewUserRepository(db)
	userServiceImpl := user.NewUserService(userRepositoryImpl)
	userHandlerImpl := user.NewHandlerUser(userServiceImpl)
	engine := app.SetupRoutes(userHandlerImpl)
	return engine
}

// injector.go:

var userSet = wire.NewSet(user.NewUserRepository, wire.Bind(new(user.UserRepository), new(*user.UserRepositoryImpl)), user.NewUserService, wire.Bind(new(user.UserService), new(*user.UserServiceImpl)), user.NewHandlerUser, wire.Bind(new(user.UserHandler), new(*user.UserHandlerImpl)))

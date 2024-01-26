//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	apiAuth "go-gin-api/api/auth"
	"go-gin-api/domain/auth"
	"go-gin-api/infrastructure"
)

var userSet = wire.NewSet(
	auth.NewUserRepository,
	wire.Bind(new(auth.Repository), new(*auth.RepositoryImpl)),
	auth.NewAuthService,
	wire.Bind(new(auth.Service), new(*auth.ServiceImpl)),
	apiAuth.NewHandlerAuth,
	wire.Bind(new(apiAuth.Handler), new(*apiAuth.HandlerImpl)),
)

func InitializedApp() *gin.Engine {
	wire.Build(
		infrastructure.ConnectToDatabase,
		userSet,
		infrastructure.SetupRoutes,
	)
	return nil
}

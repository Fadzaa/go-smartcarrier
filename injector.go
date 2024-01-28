//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	apiAuth "go-gin-api/api/auth"
	apiCourse "go-gin-api/api/course"
	apiJob "go-gin-api/api/job"
	"go-gin-api/domain/auth"
	"go-gin-api/domain/course"
	"go-gin-api/domain/job"

	"go-gin-api/infrastructure"
)

var userSet = wire.NewSet(
	auth.NewUserRepository,
	wire.Bind(new(auth.Repository), new(*auth.RepositoryImpl)),
	auth.NewAuthService,
	wire.Bind(new(auth.Service), new(*auth.ServiceImpl)),
	apiAuth.NewAuthHandler,
	wire.Bind(new(apiAuth.Handler), new(*apiAuth.HandlerImpl)),
)

var jobSet = wire.NewSet(
	job.NewJobRepository,
	wire.Bind(new(job.Repository), new(*job.RepositoryImpl)),
	job.NewJobService,
	wire.Bind(new(job.Service), new(*job.ServiceImpl)),
	apiJob.NewJobHandler,
	wire.Bind(new(apiJob.Handler), new(*apiJob.HandlerImpl)),
)

var courseSet = wire.NewSet(
	course.NewCourseRepository,
	wire.Bind(new(course.Repository), new(*course.RepositoryImpl)),
	course.NewCourseService,
	wire.Bind(new(course.Service), new(*course.ServiceImpl)),
	apiCourse.NewCourseHandler,
	wire.Bind(new(apiCourse.Handler), new(*apiCourse.HandlerImpl)),
)

func InitializedApp() *gin.Engine {
	wire.Build(
		infrastructure.ConnectToDatabase,
		userSet,
		jobSet,
		courseSet,
		infrastructure.SetupRoutes,
	)
	return nil
}

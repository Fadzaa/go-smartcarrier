// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	auth2 "go-gin-api/api/auth"
	course2 "go-gin-api/api/course"
	job2 "go-gin-api/api/job"
	"go-gin-api/domain/auth"
	"go-gin-api/domain/course"
	"go-gin-api/domain/job"
	"go-gin-api/infrastructure"
)

// Injectors from injector.go:

func InitializedApp() *gin.Engine {
	db := infrastructure.ConnectToDatabase()
	repositoryImpl := auth.NewUserRepository(db)
	serviceImpl := auth.NewAuthService(repositoryImpl)
	handlerImpl := auth2.NewAuthHandler(serviceImpl)
	jobRepositoryImpl := job.NewJobRepository(db)
	jobServiceImpl := job.NewJobService(jobRepositoryImpl)
	jobHandlerImpl := job2.NewJobHandler(jobServiceImpl)
	courseRepositoryImpl := course.NewCourseRepository(db)
	courseServiceImpl := course.NewCourseService(courseRepositoryImpl)
	courseHandlerImpl := course2.NewCourseHandler(courseServiceImpl)
	engine := infrastructure.SetupRoutes(handlerImpl, jobHandlerImpl, courseHandlerImpl)
	return engine
}

// injector.go:

var userSet = wire.NewSet(auth.NewUserRepository, wire.Bind(new(auth.Repository), new(*auth.RepositoryImpl)), auth.NewAuthService, wire.Bind(new(auth.Service), new(*auth.ServiceImpl)), auth2.NewAuthHandler, wire.Bind(new(auth2.Handler), new(*auth2.HandlerImpl)))

var jobSet = wire.NewSet(job.NewJobRepository, wire.Bind(new(job.Repository), new(*job.RepositoryImpl)), job.NewJobService, wire.Bind(new(job.Service), new(*job.ServiceImpl)), job2.NewJobHandler, wire.Bind(new(job2.Handler), new(*job2.HandlerImpl)))

var courseSet = wire.NewSet(course.NewCourseRepository, wire.Bind(new(course.Repository), new(*course.RepositoryImpl)), course.NewCourseService, wire.Bind(new(course.Service), new(*course.ServiceImpl)), course2.NewCourseHandler, wire.Bind(new(course2.Handler), new(*course2.HandlerImpl)))

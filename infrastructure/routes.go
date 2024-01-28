package infrastructure

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/api/auth"
	"go-gin-api/api/course"
	"go-gin-api/api/job"
)

func SetupRoutes(userHandler auth.Handler, jobHandler job.Handler, courseHandler course.Handler) *gin.Engine {
	router := gin.Default()

	a := router.Group("/auth")
	a.POST("/register", userHandler.RegisterUserHandler)
	a.POST("/login", userHandler.LoginUserHandler)

	j := router.Group("/job")
	j.GET("/", jobHandler.GetAllJob)
	j.GET("/:id", jobHandler.GetJobByID)
	j.POST("/", jobHandler.CreateJob)
	j.PUT("/:id", jobHandler.UpdateJob)
	j.DELETE("/:id", jobHandler.DeleteJob)

	c := router.Group("/course")
	c.GET("/", courseHandler.GetAllCourse)
	c.GET("/:id", courseHandler.GetCourseByID)
	c.POST("/", courseHandler.CreateCourse)
	c.PUT("/:id", courseHandler.UpdateCourse)
	c.DELETE("/:id", courseHandler.DeleteCourse)

	return router
}

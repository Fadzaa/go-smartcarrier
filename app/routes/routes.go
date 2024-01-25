package routes

import (
	"github.com/gin-gonic/gin"
	userHandler "go-gin-api/app/handlers/user"
	"go-gin-api/domains/user"

	"go-gin-api/infrastructure"
)

func SetupRoutes() {
	db := infrastructure.ConnectToDatabase()
	//defer db.DB().

	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	handlerUser := userHandler.NewHandlerUser(userService)

	//jobRepo := job.NewRepositoryJob(db)
	//jobService := job.NewServiceJob(jobRepo)
	//job.NewHandlerJob(e, jobService)

	router := gin.Default()

	u := router.Group("/user")
	u.GET("/", handlerUser.GetAllUserHandler)
	u.GET("/:id", handlerUser.GetUserByIDHandler)
	u.POST("/", handlerUser.CreateUserHandler)
	u.PUT("/:id", handlerUser.UpdateUserHandler)
	u.DELETE("/:id", handlerUser.DeleteUserHandler)

	//j := router.Group("/job")
	//j.GET("/", handlerJob.GetAllJobHandler)
	//j.GET("/:id", handlerJob.GetJobByIDHandler)
	//j.POST("/", handlerJob.CreateJobHandler)
	//j.PUT("/:id", handlerJob.UpdateJobHandler)
	//j.DELETE("/:id", handlerJob.DeleteJobHandler)

}

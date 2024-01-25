package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-api/domains/job"
	"go-gin-api/domains/user"
	coffeeHandler "go-gin-api/handlers/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//Open Database Connectoin
	dsn := "root:@tcp(127.0.0.1:3306)/api_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&user.User{}, &job.Job{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Database Connected")

	userRepository := user.NewUserRepository(db)
	users, err := userRepository.FindAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(users)

	user, err := userRepository.FindUserByID(2)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)

	//createUser, err := userRepository.CreateUser(user.User{
	//	Name:        "Syahran Fadhil",
	//	Email:       "syahran@mail.com",
	//	Password:    "syahran123",
	//	PhoneNumber: "08123456789",
	//	Role:        "Software Engineer",
	//	Gender:      1,
	//	Province:    "Jawa Timur",
	//	City:        "Semarang",
	//	CreatedAt:   time.Now(),
	//	UpdatedAt:   time.Now(),
	//})
	//
	//fmt.Println(createUser)

	fmt.Println("User Repository")
	//userService := user.NewUserService(userRepository)
	//Initialize Gin Router
	router := gin.Default()
	router.GET("/", coffeeHandler.RootHandler)

	c := router.Group("/user")
	c.GET("/", coffeeHandler.GetCoffeeHandler)
	c.POST("/", coffeeHandler.CreateCoffeeHandler)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}

}

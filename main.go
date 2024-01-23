package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-api/handler/coffee"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/coffee_api?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Database Connected")

	router := gin.Default()

	router.GET("/", coffee.RootHandler)

	c := router.Group("/coffee")
	c.GET("/", coffee.GetCoffeeHandler)
	c.POST("/", coffee.CreateCoffeeHandler)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}

}

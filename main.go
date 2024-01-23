package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/handler"
)

func main() {
	router := gin.Default()

	router.GET("/", handler.RootHandler)

	coffee := router.Group("/coffee")
	coffee.GET("/", handler.GetCoffeeHandler)
	coffee.POST("/", handler.CreateCoffeeHandler)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}

}

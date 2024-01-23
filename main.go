package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/handler/coffee"
)

func main() {
	router := gin.Default()

	router.GET("/", coffee.RootHandler)

	c := router.Group("/coffee")
	c.GET("/", coffee.GetCoffeeHandler)
	c.POST("/", coffee.CreateCoffeeHandler)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}

}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)

	coffee := router.Group("/coffee")
	coffee.GET("/", getCoffeeHandler)
	coffee.POST("/", createCoffeeHandler)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}

}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":    "Hello There Welcome to Coffee API!",
		"created_by": "Fattah Anggit",
	})
}

func getCoffeeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":   "Ethiopian Yirgacheffe Cheffe",
		"origin": "Ethiopia",
	})
}

type Coffee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Origin string `json:"origin"`
}

func createCoffeeHandler(c *gin.Context) {
	var coffee Coffee
	err := c.ShouldBindJSON(&coffee)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Coffee Successfully Created",
		"name":    coffee.Name,
		"origin":  coffee.Origin,
	})
}

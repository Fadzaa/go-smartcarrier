package handler

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/model"
	"net/http"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":    "Hello There Welcome to Coffee API!",
		"created_by": "Fattah Anggit",
	})
}

func GetCoffeeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":   "Ethiopian Yirgacheffe Cheffe",
		"origin": "Ethiopia",
	})
}

func CreateCoffeeHandler(c *gin.Context) {
	var coffee model.Coffee
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

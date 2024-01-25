package user

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/domains/user"
	"net/http"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":    "Hello There Welcome to User API!",
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
	var users user.User
	err := c.ShouldBindJSON(&users)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Successfully Created",
		"name":    users.Name,
		"origin":  users.Email,
	})
}

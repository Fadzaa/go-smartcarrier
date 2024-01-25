package user

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/domains/user"
	"net/http"
	"strconv"
)

type HandlerUser interface {
	GetAllUserHandler(c *gin.Context)
	GetUserByIDHandler(c *gin.Context)
	CreateUserHandler(c *gin.Context)
	UpdateUserHandler(c *gin.Context)
	DeleteUserHandler(c *gin.Context)
}

type handlerUser struct {
	serviceUser user.ServiceUser
}

func NewHandlerUser(serviceUser user.ServiceUser) *handlerUser {
	return &handlerUser{serviceUser}
}

func RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello There Welcome to Dinacom API SmartCarrier!",
	})
}

func (h *handlerUser) GetAllUserHandler(c *gin.Context) {
	users, err := h.serviceUser.GetAllUser()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": users,
	})

}

func (h *handlerUser) GetUserByIDHandler(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	userByID, err := h.serviceUser.GetUserByID(idInt)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": userByID,
	})
}

func (h *handlerUser) CreateUserHandler(c *gin.Context) {
	var user user.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	createUser, err := h.serviceUser.CreateUser(user)

	c.JSON(200, gin.H{
		"message": "User Successfully Created",
		"data":    createUser,
	})
}

func (h *handlerUser) UpdateUserHandler(c *gin.Context) {
	var user user.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	updateUser, err := h.serviceUser.UpdateUser(user)

	c.JSON(200, gin.H{
		"message": "User Successfully Updated",
		"data":    updateUser,
	})
}

func (h *handlerUser) DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	deleteUser, err := h.serviceUser.DeleteUser(idInt)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User Successfully Deleted",
		"data":    deleteUser,
	})
}

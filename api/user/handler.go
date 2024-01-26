package user

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/domain/user"
	"net/http"
	"strconv"
)

type UserHandler interface {
	GetAllUserHandler(c *gin.Context)
	GetUserByIDHandler(c *gin.Context)
	CreateUserHandler(c *gin.Context)
	UpdateUserHandler(c *gin.Context)
	DeleteUserHandler(c *gin.Context)
}

type UserHandlerImpl struct {
	serviceUser user.UserService
}

func NewHandlerUser(serviceUser user.UserService) *UserHandlerImpl {
	return &UserHandlerImpl{serviceUser}
}

func RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello There Welcome to Dinacom API SmartCarrier!",
	})
}

func (h *UserHandlerImpl) GetAllUserHandler(c *gin.Context) {
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

func (h *UserHandlerImpl) GetUserByIDHandler(c *gin.Context) {
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

func (h *UserHandlerImpl) CreateUserHandler(c *gin.Context) {
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

func (h *UserHandlerImpl) UpdateUserHandler(c *gin.Context) {
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

func (h *UserHandlerImpl) DeleteUserHandler(c *gin.Context) {
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

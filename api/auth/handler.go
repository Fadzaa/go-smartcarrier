package auth

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/domain/auth"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Handler interface {
	RegisterUserHandler(c *gin.Context)
	LoginUserHandler(c *gin.Context)
}

type HandlerImpl struct {
	serviceUser auth.Service
}

func NewAuthHandler(serviceUser auth.Service) *HandlerImpl {
	return &HandlerImpl{serviceUser}
}

func (h *HandlerImpl) RegisterUserHandler(c *gin.Context) {
	var user auth.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error Binding JSON",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error Hashing Password",
		})
		return
	}

	user.Password = string(hashedPassword)

	createUser, err := h.serviceUser.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error Creating User",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User Successfully Created",
		"data":    createUser,
	})
}

func (h *HandlerImpl) LoginUserHandler(c *gin.Context) {
	var user auth.User

	//Body
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error Binding JSON",
		})
		return
	}

	//Result
	userLogin, err := h.serviceUser.LoginUser(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error Login User",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(user.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email and Password ",
		})
		return
	}

	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//	"sub": user.ID,
	//	"exp": time.Now().Add(time.Hour).Unix(),
	//})

	//tokenString, err := token.SignedString([]byte("secretjfkladsjflksjaflaejfiejf"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error Creating Token",
		})
		return
	}

	//if i want to set cookie
	//c.SetSameSite(http.SameSiteLaxMode)
	//c.SetCookie("Authorization", tokenString, 3600, "/", "localhost", false, true)

	//if i want to set token in header
	//c.JSON(200, gin.H{
	//	"message": "User Successfully Login",
	//	"data":    userLogin,
	//	"token":   tokenString,
	//})

	c.JSON(200, userLogin)
}

package course

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/domain/course"
	"net/http"
	"strconv"
)

type Handler interface {
	GetAllCourse(c *gin.Context)
	GetCourseByID(c *gin.Context)
	CreateCourse(c *gin.Context)
	UpdateCourse(c *gin.Context)
	DeleteCourse(c *gin.Context)
}

type HandlerImpl struct {
	serviceCourse course.Service
}

func NewCourseHandler(serviceCourse course.Service) *HandlerImpl {
	return &HandlerImpl{serviceCourse}
}

func (h *HandlerImpl) GetAllCourse(c *gin.Context) {
	courses, err := h.serviceCourse.GetAllCourse()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	//c.JSON(http.StatusOK, gin.H{
	//	"data": courses,
	//})

	c.JSON(http.StatusOK, courses)
}

func (h *HandlerImpl) GetCourseByID(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	courseByID, err := h.serviceCourse.GetCourseByID(idInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": courseByID,
	})

}

func (h *HandlerImpl) CreateCourse(c *gin.Context) {
	var courseBind course.Course
	err := c.ShouldBindJSON(&courseBind)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error Binding JSON",
		})
		return
	}

	createCourse, err := h.serviceCourse.CreateCourse(courseBind)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error Creating Course",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Course Successfully Created",
		"data":    createCourse,
	})
}

func (h *HandlerImpl) UpdateCourse(c *gin.Context) {
	var courseBind course.Course
	err := c.ShouldBindJSON(&courseBind)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error Binding JSON",
		})
		return
	}

	updateCourse, err := h.serviceCourse.UpdateCourse(courseBind)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error Updating Course",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Course Successfully Updated",
		"data":    updateCourse,
	})
}

func (h *HandlerImpl) DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	_, err := h.serviceCourse.DeleteCourse(idInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error Deleting Course",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Course Successfully Deleted",
	})
}

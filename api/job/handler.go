package job

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/domain/job"
	"net/http"
	"strconv"
)

type Handler interface {
	GetAllJob(c *gin.Context)
	GetJobByID(c *gin.Context)
	CreateJob(c *gin.Context)
	UpdateJob(c *gin.Context)
	DeleteJob(c *gin.Context)
}

type HandlerImpl struct {
	serviceJob job.Service
}

func NewJobHandler(serviceJob job.Service) *HandlerImpl {
	return &HandlerImpl{serviceJob}
}

func (h *HandlerImpl) GetAllJob(c *gin.Context) {
	jobs, err := h.serviceJob.GetAllJob()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": jobs,
	})
}

func (h *HandlerImpl) GetJobByID(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	jobByID, err := h.serviceJob.GetJobByID(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": jobByID,
	})

}

func (h *HandlerImpl) CreateJob(c *gin.Context) {
	var jobBind job.Job
	err := c.ShouldBindJSON(&jobBind)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error Binding JSON",
		})
		return
	}

	createJob, err := h.serviceJob.CreateJob(jobBind)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error Creating Job",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Job Successfully Created",
		"data":    createJob,
	})
}

func (h *HandlerImpl) UpdateJob(c *gin.Context) {
	var jobBind job.Job
	err := c.ShouldBindJSON(&jobBind)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error Binding JSON",
		})
		return
	}

	updateJob, err := h.serviceJob.UpdateJob(jobBind)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error Updating Job",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Job Successfully Updated",
		"data":    updateJob,
	})
}

func (h *HandlerImpl) DeleteJob(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	_, err := h.serviceJob.DeleteJob(idInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error Deleting Job",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Course Successfully Job",
	})
}

package server

import (
	"net/http"

	"github.com/bersennaidoo/agentco/domain/core/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h Handler) TestHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Hello From Gin"})
}

func (h Handler) ModifyJobApplicationWithId(c *gin.Context, id string) {

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h Handler) ListOrSearchAvailableJobs(c *gin.Context, params model.ListOrSearchAvailableJobsParams) {

	c.JSON(http.StatusOK, gin.H{"params": params})
}

func (h Handler) CreateJob(c *gin.Context) {

	var j model.Job

	if err := c.ShouldBindJSON(&j); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Job creation successful."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Job creation failed!",
			"error":   err.Error(),
		})
	}
}

func (h Handler) DeleteJobWithId(c *gin.Context, id string) {

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h Handler) ViewJobWithId(c *gin.Context, id string) {

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h Handler) ModifyJobWithId(c *gin.Context, id string) {

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h Handler) ViewApplicationsForJob(c *gin.Context, id string) {

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h Handler) CreateJobApplication(c *gin.Context, id string) {

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h Handler) StartSession(c *gin.Context) {

	var s model.Session
	if err := c.ShouldBindJSON(&s); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Session creation successful."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Session creation failed!",
			"error":   err.Error(),
		})
	}
}

func (h Handler) RegisterUser(c *gin.Context) {

	var u model.User
	if err := c.ShouldBindJSON(&u); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "User creation successful."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User creation failed!",
			"error":   err.Error(),
		})
	}
}

func (h Handler) DeleteUserWithId(c *gin.Context, id string) {

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h Handler) ViewUserWithId(c *gin.Context, id string) {

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h Handler) ModifyUserWithId(c *gin.Context, id string) {

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h Handler) ViewApplicationsForUser(c *gin.Context, id int) {

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h Handler) ListJobsForUser(c *gin.Context, id string) {

	c.JSON(http.StatusOK, gin.H{"id": id})
}

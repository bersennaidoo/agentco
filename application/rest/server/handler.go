package server

import (
	"net/http"

	"github.com/bersennaidoo/agentco/domain/core/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h Handler) ModifyJobApplicationWithId(c *gin.Context, id string) {

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h Handler) ListOrSearchAvailableJobs(c *gin.Context, params model.ListOrSearchAvailableJobsParams) {

	c.JSON(http.StatusOK, gin.H{"params": params})
}

func (h Handler) CreateJob(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"id": "CreateJob"})
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

	c.JSON(http.StatusOK, gin.H{"id": "StartSession"})
}

func (h Handler) RegisterUser(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"id": "RegisterUser"})
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

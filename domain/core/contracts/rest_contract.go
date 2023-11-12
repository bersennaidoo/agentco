package contracts

import (
	"github.com/bersennaidoo/agentco/domain/core/model"
	"github.com/gin-gonic/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	TestHandler(c *gin.Context)
	// Modify Job Application
	// (PUT /job-applications/{id})
	ModifyJobApplicationWithId(c *gin.Context, id string)
	// List/Search Available Jobs
	// (GET /jobs)
	ListOrSearchAvailableJobs(c *gin.Context, params model.ListOrSearchAvailableJobsParams)
	// Create Job
	// (POST /jobs)
	CreateJob(c *gin.Context)
	// Delete Job
	// (DELETE /jobs/{id})
	DeleteJobWithId(c *gin.Context, id string)
	// View Job
	// (GET /jobs/{id})
	ViewJobWithId(c *gin.Context, id string)
	// Modify Job
	// (PUT /jobs/{id})
	ModifyJobWithId(c *gin.Context, id string)
	// List Applications For Job
	// (GET /jobs/{id}/job-applications)
	ViewApplicationsForJob(c *gin.Context, id string)
	// Create Job Application
	// (POST /jobs/{id}/job-applications)
	CreateJobApplication(c *gin.Context, id string)
	// Start Session (Login)
	// (POST /sessions)
	StartSession(c *gin.Context)
	// Register User
	// (POST /users)
	RegisterUser(c *gin.Context)
	// Delete User
	// (DELETE /users/{id})
	DeleteUserWithId(c *gin.Context, id string)
	// View User
	// (GET /users/{id})
	ViewUserWithId(c *gin.Context, id string)
	// Modify User
	// (PUT /users/{id})
	ModifyUserWithId(c *gin.Context, id string)
	// List Applications For User
	// (GET /users/{id}/job-applications)
	ViewApplicationsForUser(c *gin.Context, id int)
	// List Jobs For User
	// (GET /users/{id}/jobs)
	ListJobsForUser(c *gin.Context, id string)
}

package handlers

import (
	"fmt"
	"net/http"

	"github.com/bersennaidoo/agentco/domain/models"
	"github.com/bersennaidoo/agentco/infrastructure/repositories/mongo"
)

type Handler struct {
	userRepository *mongo.UserRepository
}

func New(userRepository *mongo.UserRepository) *Handler {
	return &Handler{
		userRepository: userRepository,
	}
}

func (h *Handler) DeleteJobApplication(w http.ResponseWriter, r *http.Request, id string) {
}

func (h *Handler) UpdateJobApplication(w http.ResponseWriter, r *http.Request, id string) {
}

func (h *Handler) GetJobs(w http.ResponseWriter, r *http.Request, params models.GetJobsParams) {
}

func (h *Handler) PostJobs(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) DeleteJobsId(w http.ResponseWriter, r *http.Request, id string) {
}

func (h *Handler) GetJobsId(w http.ResponseWriter, r *http.Request, id string) {
}

func (h *Handler) PutJobsId(w http.ResponseWriter, r *http.Request, id string) {
}

func (h *Handler) GetApplicationsByJobId(w http.ResponseWriter, r *http.Request, id string) {
}

func (h *Handler) CreateJobApplication(w http.ResponseWriter, r *http.Request, id string) {
}

func (h *Handler) StartSession(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) PostUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from PostUsers")
}

func (h *Handler) DeleteUsersId(w http.ResponseWriter, r *http.Request, id string) {
}

func (h *Handler) GetUsersId(w http.ResponseWriter, r *http.Request, id string) {
}

func (h *Handler) PutUsersId(w http.ResponseWriter, r *http.Request, id string) {
}

func (h *Handler) GetJobApplicationsForUser(w http.ResponseWriter, r *http.Request, id string) {
}

func (h *Handler) GetJobsForUser(w http.ResponseWriter, r *http.Request, id string) {
}

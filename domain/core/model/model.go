package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	SessionTokenScopes = "SessionToken.Scopes"
)

// Defines values for JobApplicationStatus.
const (
	Applying JobApplicationStatus = "applying"
	Approved JobApplicationStatus = "approved"
	Rejected JobApplicationStatus = "rejected"
)

// Defines values for UserRoles.
const (
	Admin     UserRoles = "Admin"
	PetOwner  UserRoles = "PetOwner"
	PetSitter UserRoles = "PetSitter"
)

// Cat defines model for Cat.
type Cat struct {
	Breed   *string `json:"breed,omitempty"`
	Species string  `json:"species"`
}

// Dog defines model for Dog.
type Dog struct {
	Breed   *string `json:"breed,omitempty"`
	Size    *string `json:"size,omitempty"`
	Species string  `json:"species"`
}

// Job defines model for Job.
type Job struct {
	Activity      *string `json:"activity,omitempty"`
	CreatorUserId *int    `json:"creator_user_id,omitempty"`
	Dog           *struct {
		Breed   *string `json:"breed,omitempty"`
		Size    *string `json:"size,omitempty"`
		Species string  `json:"species"`
	} `json:"dog,omitempty"`
	EndTime   *time.Time `json:"end_time,omitempty"`
	Id        *int       `json:"id,omitempty"`
	Pets      *[]Pet     `json:"pets,omitempty"`
	StartTime *time.Time `json:"start_time,omitempty"`
}

// JobApplication defines model for JobApplication.
type JobApplication struct {
	Id     *string               `json:"id,omitempty"`
	JobId  *string               `json:"job_id,omitempty"`
	Status *JobApplicationStatus `json:"status,omitempty"`
	UserId *string               `json:"user_id,omitempty"`
}

// JobApplicationStatus defines model for JobApplication.Status.
type JobApplicationStatus string

// OASError defines model for OASError.
type OASError struct {
	Errors *[]struct {
		// ErrorCode Code indicating error type
		ErrorCode *string `json:"errorCode,omitempty"`

		// Message Human-readable error message
		Message *string `json:"message,omitempty"`

		// Path For input validation errors, identifies where in the JSON request body the error occured
		Path *string `json:"path,omitempty"`
	} `json:"errors,omitempty"`

	// Message Human-readable error message
	Message *string `json:"message,omitempty"`
}

// Pet defines model for Pet.
type Pet struct {
	Age   *int    `json:"age,omitempty"`
	Name  *string `json:"name,omitempty"`
	union json.RawMessage
}

// Problem defines model for Problem.
type Problem struct {
	// Detail Human-readable error details
	Detail *string `json:"detail,omitempty"`

	// Instance URI indicating error instance
	Instance *string `json:"instance,omitempty"`

	// Status HTTP status code
	Status *int `json:"status,omitempty"`

	// Title Human-readable error title
	Title *string `json:"title,omitempty"`

	// Type URI indicating error type
	Type *string `json:"type,omitempty"`
}

// Session defines model for Session.
type Session struct {
	AuthHeader *string `json:"auth_header,omitempty"`
	UserId     *string `json:"user_id,omitempty"`
}

// User defines model for User.
type User struct {
	Email    openapi_types.Email `json:"email"`
	FullName string              `json:"full_name"`
	Id       *int                `json:"id,omitempty"`
	Password *string             `json:"password,omitempty"`
	Roles    []UserRoles         `json:"roles"`
}

// UserRoles defines model for User.Roles.
type UserRoles string

// ListOrSearchAvailableJobsParams defines parameters for ListOrSearchAvailableJobs.
type ListOrSearchAvailableJobsParams struct {
	// StartTimeBefore Search jobs starting before this date and time.
	StartTimeBefore *time.Time `form:"start_time_before,omitempty" json:"start_time_before,omitempty"`

	// StartTimeAfter Search jobs starting after this date and time.
	StartTimeAfter *time.Time `form:"start_time_after,omitempty" json:"start_time_after,omitempty"`

	// EndTimeBefore Search jobs ending before this date and time.
	EndTimeBefore *time.Time `form:"end_time_before,omitempty" json:"end_time_before,omitempty"`

	// EndTimeAfter Search jobs ending after this date and time.
	EndTimeAfter *time.Time `form:"end_time_after,omitempty" json:"end_time_after,omitempty"`

	// Activity Performs a full-text search for the phrase entered in job activities.
	Activity *string `form:"activity,omitempty" json:"activity,omitempty"`

	// Pets Searches for pets matching specific criteria.
	Pets *struct {
		// AgeAbove Return only pets with this age or older.
		AgeAbove *int `json:"age_above,omitempty"`

		// AgeBelow Return only pets with this age or younger.
		AgeBelow *int `json:"age_below,omitempty"`

		// Species Return only pets with this species. Provide multiple species as comma-separated values.
		Species *string `json:"species,omitempty"`
	} `json:"pets,omitempty"`

	// Limit The maximum number of results to return.
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`

	// Cursor Use the cursor from the response to access more results.
	Cursor *string `form:"cursor,omitempty" json:"cursor,omitempty"`

	// Sort Indicate the sorting key and direction for the results.
	// Use the field name, suffixed with ":asc" for ascending
	// or ":desc" for descending order.
	// Valid fields: start_time, end_time
	Sort *string `form:"sort,omitempty" json:"sort,omitempty"`
}

// StartSessionJSONBody defines parameters for StartSession.
type StartSessionJSONBody struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

// ModifyJobApplicationWithIdJSONRequestBody defines body for ModifyJobApplicationWithId for application/json ContentType.
type ModifyJobApplicationWithIdJSONRequestBody = JobApplication

// ModifyJobWithIdJSONRequestBody defines body for ModifyJobWithId for application/json ContentType.
type ModifyJobWithIdJSONRequestBody = Job

// CreateJobApplicationJSONRequestBody defines body for CreateJobApplication for application/json ContentType.
type CreateJobApplicationJSONRequestBody = JobApplication

// StartSessionJSONRequestBody defines body for StartSession for application/json ContentType.
type StartSessionJSONRequestBody StartSessionJSONBody

// RegisterUserJSONRequestBody defines body for RegisterUser for application/json ContentType.
type RegisterUserJSONRequestBody = User

// ModifyUserWithIdJSONRequestBody defines body for ModifyUserWithId for application/json ContentType.
type ModifyUserWithIdJSONRequestBody = User

// AsCat returns the union data inside the Pet as a Cat
func (t Pet) AsCat() (Cat, error) {
	var body Cat
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCat overwrites any union data inside the Pet as the provided Cat
func (t *Pet) FromCat(v Cat) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCat performs a merge with any union data inside the Pet, using the provided Cat
func (t *Pet) MergeCat(v Cat) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsDog returns the union data inside the Pet as a Dog
func (t Pet) AsDog() (Dog, error) {
	var body Dog
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDog overwrites any union data inside the Pet as the provided Dog
func (t *Pet) FromDog(v Dog) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDog performs a merge with any union data inside the Pet, using the provided Dog
func (t *Pet) MergeDog(v Dog) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t Pet) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	if err != nil {
		return nil, err
	}
	object := make(map[string]json.RawMessage)
	if t.union != nil {
		err = json.Unmarshal(b, &object)
		if err != nil {
			return nil, err
		}
	}

	if t.Age != nil {
		object["age"], err = json.Marshal(t.Age)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'age': %w", err)
		}
	}

	if t.Name != nil {
		object["name"], err = json.Marshal(t.Name)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'name': %w", err)
		}
	}
	b, err = json.Marshal(object)
	return b, err
}

func (t *Pet) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	if err != nil {
		return err
	}
	object := make(map[string]json.RawMessage)
	err = json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["age"]; found {
		err = json.Unmarshal(raw, &t.Age)
		if err != nil {
			return fmt.Errorf("error reading 'age': %w", err)
		}
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &t.Name)
		if err != nil {
			return fmt.Errorf("error reading 'name': %w", err)
		}
	}

	return err
}

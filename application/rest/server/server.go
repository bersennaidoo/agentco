package server

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/bersennaidoo/agentco/domain/core/model"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
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

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// ModifyJobApplicationWithId operation middleware
func (siw *ServerInterfaceWrapper) ModifyJobApplicationWithId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(model.SessionTokenScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ModifyJobApplicationWithId(c, id)
}

// ListOrSearchAvailableJobs operation middleware
func (siw *ServerInterfaceWrapper) ListOrSearchAvailableJobs(c *gin.Context) {

	var err error

	c.Set(model.SessionTokenScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params model.ListOrSearchAvailableJobsParams

	// ------------- Optional query parameter "start_time_before" -------------

	err = runtime.BindQueryParameter("form", true, false, "start_time_before", c.Request.URL.Query(), &params.StartTimeBefore)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter start_time_before: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "start_time_after" -------------

	err = runtime.BindQueryParameter("form", true, false, "start_time_after", c.Request.URL.Query(), &params.StartTimeAfter)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter start_time_after: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "end_time_before" -------------

	err = runtime.BindQueryParameter("form", true, false, "end_time_before", c.Request.URL.Query(), &params.EndTimeBefore)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter end_time_before: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "end_time_after" -------------

	err = runtime.BindQueryParameter("form", true, false, "end_time_after", c.Request.URL.Query(), &params.EndTimeAfter)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter end_time_after: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "activity" -------------

	err = runtime.BindQueryParameter("form", true, false, "activity", c.Request.URL.Query(), &params.Activity)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter activity: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "pets" -------------

	err = runtime.BindQueryParameter("deepObject", true, false, "pets", c.Request.URL.Query(), &params.Pets)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter pets: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", c.Request.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter limit: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "cursor" -------------

	err = runtime.BindQueryParameter("form", true, false, "cursor", c.Request.URL.Query(), &params.Cursor)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter cursor: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "sort" -------------

	err = runtime.BindQueryParameter("form", true, false, "sort", c.Request.URL.Query(), &params.Sort)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter sort: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ListOrSearchAvailableJobs(c, params)
}

// CreateJob operation middleware
func (siw *ServerInterfaceWrapper) CreateJob(c *gin.Context) {

	c.Set(model.SessionTokenScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateJob(c)
}

// DeleteJobWithId operation middleware
func (siw *ServerInterfaceWrapper) DeleteJobWithId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(model.SessionTokenScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteJobWithId(c, id)
}

// ViewJobWithId operation middleware
func (siw *ServerInterfaceWrapper) ViewJobWithId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(model.SessionTokenScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ViewJobWithId(c, id)
}

// ModifyJobWithId operation middleware
func (siw *ServerInterfaceWrapper) ModifyJobWithId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(model.SessionTokenScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ModifyJobWithId(c, id)
}

// ViewApplicationsForJob operation middleware
func (siw *ServerInterfaceWrapper) ViewApplicationsForJob(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(model.SessionTokenScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ViewApplicationsForJob(c, id)
}

// CreateJobApplication operation middleware
func (siw *ServerInterfaceWrapper) CreateJobApplication(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(model.SessionTokenScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateJobApplication(c, id)
}

// StartSession operation middleware
func (siw *ServerInterfaceWrapper) StartSession(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.StartSession(c)
}

// RegisterUser operation middleware
func (siw *ServerInterfaceWrapper) RegisterUser(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.RegisterUser(c)
}

// DeleteUserWithId operation middleware
func (siw *ServerInterfaceWrapper) DeleteUserWithId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(model.SessionTokenScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteUserWithId(c, id)
}

// ViewUserWithId operation middleware
func (siw *ServerInterfaceWrapper) ViewUserWithId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(model.SessionTokenScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ViewUserWithId(c, id)
}

// ModifyUserWithId operation middleware
func (siw *ServerInterfaceWrapper) ModifyUserWithId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(model.SessionTokenScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ModifyUserWithId(c, id)
}

// ViewApplicationsForUser operation middleware
func (siw *ServerInterfaceWrapper) ViewApplicationsForUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(model.SessionTokenScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ViewApplicationsForUser(c, id)
}

// ListJobsForUser operation middleware
func (siw *ServerInterfaceWrapper) ListJobsForUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(model.SessionTokenScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ListJobsForUser(c, id)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.PUT(options.BaseURL+"/job-applications/:id", wrapper.ModifyJobApplicationWithId)
	router.GET(options.BaseURL+"/jobs", wrapper.ListOrSearchAvailableJobs)
	router.POST(options.BaseURL+"/jobs", wrapper.CreateJob)
	router.DELETE(options.BaseURL+"/jobs/:id", wrapper.DeleteJobWithId)
	router.GET(options.BaseURL+"/jobs/:id", wrapper.ViewJobWithId)
	router.PUT(options.BaseURL+"/jobs/:id", wrapper.ModifyJobWithId)
	router.GET(options.BaseURL+"/jobs/:id/job-applications", wrapper.ViewApplicationsForJob)
	router.POST(options.BaseURL+"/jobs/:id/job-applications", wrapper.CreateJobApplication)
	router.POST(options.BaseURL+"/sessions", wrapper.StartSession)
	router.POST(options.BaseURL+"/users", wrapper.RegisterUser)
	router.DELETE(options.BaseURL+"/users/:id", wrapper.DeleteUserWithId)
	router.GET(options.BaseURL+"/users/:id", wrapper.ViewUserWithId)
	router.PUT(options.BaseURL+"/users/:id", wrapper.ModifyUserWithId)
	router.GET(options.BaseURL+"/users/:id/job-applications", wrapper.ViewApplicationsForUser)
	router.GET(options.BaseURL+"/users/:id/jobs", wrapper.ListJobsForUser)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"OV3+LwAA//9Dl6b6OCcAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

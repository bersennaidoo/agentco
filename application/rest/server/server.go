package server

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/bersennaidoo/agentco/domain/models"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Delete application
	// (DELETE /job-applications/{id})
	DeleteJobApplication(w http.ResponseWriter, r *http.Request, id string)
	// Update application details
	// (PUT /job-applications/{id})
	UpdateJobApplication(w http.ResponseWriter, r *http.Request, id string)
	// List available jobs
	// (GET /jobs)
	GetJobs(w http.ResponseWriter, r *http.Request, params models.GetJobsParams)
	// Post new Job
	// (POST /jobs)
	PostJobs(w http.ResponseWriter, r *http.Request)
	// Remove Job
	// (DELETE /jobs/{id})
	DeleteJobsId(w http.ResponseWriter, r *http.Request, id string)
	// Get Job Details
	// (GET /jobs/{id})
	GetJobsId(w http.ResponseWriter, r *http.Request, id string)
	// Update Job Details
	// (PUT /jobs/{id})
	PutJobsId(w http.ResponseWriter, r *http.Request, id string)
	// Get all applications for this job.
	// (GET /jobs/{id}/job-applications)
	GetApplicationsByJobId(w http.ResponseWriter, r *http.Request, id string)
	// Create a job application
	// (POST /jobs/{id}/job-applications)
	CreateJobApplication(w http.ResponseWriter, r *http.Request, id string)
	// Start Session (Login)
	// (POST /sessions)
	StartSession(w http.ResponseWriter, r *http.Request)
	// Register User Account
	// (POST /users)
	PostUsers(w http.ResponseWriter, r *http.Request)
	// Delete User Account
	// (DELETE /users/{id})
	DeleteUsersId(w http.ResponseWriter, r *http.Request, id string)
	// Get User Information
	// (GET /users/{id})
	GetUsersId(w http.ResponseWriter, r *http.Request, id string)
	// Update User Account
	// (PUT /users/{id})
	PutUsersId(w http.ResponseWriter, r *http.Request, id string)
	// Get a list of Job Applications that are associated with this user.
	// (GET /users/{id}/job-applications)
	GetJobApplicationsForUser(w http.ResponseWriter, r *http.Request, id string)
	// Get a list of Jobs that are associated with this user.
	// (GET /users/{id}/jobs)
	GetJobsForUser(w http.ResponseWriter, r *http.Request, id string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// DeleteJobApplication operation middleware
func (siw *ServerInterfaceWrapper) DeleteJobApplication(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", mux.Vars(r)["id"], &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, models.SessionTokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteJobApplication(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateJobApplication operation middleware
func (siw *ServerInterfaceWrapper) UpdateJobApplication(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", mux.Vars(r)["id"], &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, models.SessionTokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateJobApplication(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetJobs operation middleware
func (siw *ServerInterfaceWrapper) GetJobs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, models.SessionTokenScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params models.GetJobsParams

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", r.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "offset", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetJobs(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostJobs operation middleware
func (siw *ServerInterfaceWrapper) PostJobs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, models.SessionTokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostJobs(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteJobsId operation middleware
func (siw *ServerInterfaceWrapper) DeleteJobsId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", mux.Vars(r)["id"], &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, models.SessionTokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteJobsId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetJobsId operation middleware
func (siw *ServerInterfaceWrapper) GetJobsId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", mux.Vars(r)["id"], &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, models.SessionTokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetJobsId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutJobsId operation middleware
func (siw *ServerInterfaceWrapper) PutJobsId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", mux.Vars(r)["id"], &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, models.SessionTokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutJobsId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApplicationsByJobId operation middleware
func (siw *ServerInterfaceWrapper) GetApplicationsByJobId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", mux.Vars(r)["id"], &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, models.SessionTokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApplicationsByJobId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateJobApplication operation middleware
func (siw *ServerInterfaceWrapper) CreateJobApplication(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", mux.Vars(r)["id"], &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, models.SessionTokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateJobApplication(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// StartSession operation middleware
func (siw *ServerInterfaceWrapper) StartSession(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.StartSession(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostUsers operation middleware
func (siw *ServerInterfaceWrapper) PostUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostUsers(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteUsersId operation middleware
func (siw *ServerInterfaceWrapper) DeleteUsersId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", mux.Vars(r)["id"], &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, models.SessionTokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteUsersId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetUsersId operation middleware
func (siw *ServerInterfaceWrapper) GetUsersId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", mux.Vars(r)["id"], &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, models.SessionTokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUsersId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutUsersId operation middleware
func (siw *ServerInterfaceWrapper) PutUsersId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", mux.Vars(r)["id"], &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, models.SessionTokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutUsersId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetJobApplicationsForUser operation middleware
func (siw *ServerInterfaceWrapper) GetJobApplicationsForUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", mux.Vars(r)["id"], &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, models.SessionTokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetJobApplicationsForUser(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetJobsForUser operation middleware
func (siw *ServerInterfaceWrapper) GetJobsForUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", mux.Vars(r)["id"], &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, models.SessionTokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetJobsForUser(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{})
}

type GorillaServerOptions struct {
	BaseURL          string
	BaseRouter       *mux.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r *mux.Router) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r *mux.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options GorillaServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = mux.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.HandleFunc(options.BaseURL+"/job-applications/{id}", wrapper.DeleteJobApplication).Methods("DELETE")

	r.HandleFunc(options.BaseURL+"/job-applications/{id}", wrapper.UpdateJobApplication).Methods("PUT")

	r.HandleFunc(options.BaseURL+"/jobs", wrapper.GetJobs).Methods("GET")

	r.HandleFunc(options.BaseURL+"/jobs", wrapper.PostJobs).Methods("POST")

	r.HandleFunc(options.BaseURL+"/jobs/{id}", wrapper.DeleteJobsId).Methods("DELETE")

	r.HandleFunc(options.BaseURL+"/jobs/{id}", wrapper.GetJobsId).Methods("GET")

	r.HandleFunc(options.BaseURL+"/jobs/{id}", wrapper.PutJobsId).Methods("PUT")

	r.HandleFunc(options.BaseURL+"/jobs/{id}/job-applications", wrapper.GetApplicationsByJobId).Methods("GET")

	r.HandleFunc(options.BaseURL+"/jobs/{id}/job-applications", wrapper.CreateJobApplication).Methods("POST")

	r.HandleFunc(options.BaseURL+"/sessions", wrapper.StartSession).Methods("POST")

	r.HandleFunc(options.BaseURL+"/users", wrapper.PostUsers).Methods("POST")

	r.HandleFunc(options.BaseURL+"/users/{id}", wrapper.DeleteUsersId).Methods("DELETE")

	r.HandleFunc(options.BaseURL+"/users/{id}", wrapper.GetUsersId).Methods("GET")

	r.HandleFunc(options.BaseURL+"/users/{id}", wrapper.PutUsersId).Methods("PUT")

	r.HandleFunc(options.BaseURL+"/users/{id}/job-applications", wrapper.GetJobApplicationsForUser).Methods("GET")

	r.HandleFunc(options.BaseURL+"/users/{id}/jobs", wrapper.GetJobsForUser).Methods("GET")

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
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

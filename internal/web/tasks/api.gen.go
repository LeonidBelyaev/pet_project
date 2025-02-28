// Package tasks provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// Message defines model for Message.
type Message struct {
	Id     *uint  `json:"id,omitempty"`
	IsDone *bool  `json:"is_done,omitempty"`
	Task   string `json:"task"`

	// UserId ID of the user who owns the task
	UserId uint `json:"user_id"`
}

// PostApiTasksJSONRequestBody defines body for PostApiTasks for application/json ContentType.
type PostApiTasksJSONRequestBody = Message

// PatchApiTasksIdJSONRequestBody defines body for PatchApiTasksId for application/json ContentType.
type PatchApiTasksIdJSONRequestBody = Message

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all tasks
	// (GET /api/tasks)
	GetApiTasks(ctx echo.Context) error
	// Create a new task
	// (POST /api/tasks)
	PostApiTasks(ctx echo.Context) error
	// Delete a task by ID
	// (DELETE /api/tasks/{id})
	DeleteApiTasksId(ctx echo.Context, id uint) error
	// Update a task by ID
	// (PATCH /api/tasks/{id})
	PatchApiTasksId(ctx echo.Context, id uint) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetApiTasks converts echo context to params.
func (w *ServerInterfaceWrapper) GetApiTasks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetApiTasks(ctx)
	return err
}

// PostApiTasks converts echo context to params.
func (w *ServerInterfaceWrapper) PostApiTasks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostApiTasks(ctx)
	return err
}

// DeleteApiTasksId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteApiTasksId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteApiTasksId(ctx, id)
	return err
}

// PatchApiTasksId converts echo context to params.
func (w *ServerInterfaceWrapper) PatchApiTasksId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchApiTasksId(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/api/tasks", wrapper.GetApiTasks)
	router.POST(baseURL+"/api/tasks", wrapper.PostApiTasks)
	router.DELETE(baseURL+"/api/tasks/:id", wrapper.DeleteApiTasksId)
	router.PATCH(baseURL+"/api/tasks/:id", wrapper.PatchApiTasksId)

}

type GetApiTasksRequestObject struct {
}

type GetApiTasksResponseObject interface {
	VisitGetApiTasksResponse(w http.ResponseWriter) error
}

type GetApiTasks200JSONResponse []Message

func (response GetApiTasks200JSONResponse) VisitGetApiTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostApiTasksRequestObject struct {
	Body *PostApiTasksJSONRequestBody
}

type PostApiTasksResponseObject interface {
	VisitPostApiTasksResponse(w http.ResponseWriter) error
}

type PostApiTasks201JSONResponse Message

func (response PostApiTasks201JSONResponse) VisitPostApiTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type DeleteApiTasksIdRequestObject struct {
	Id uint `json:"id"`
}

type DeleteApiTasksIdResponseObject interface {
	VisitDeleteApiTasksIdResponse(w http.ResponseWriter) error
}

type DeleteApiTasksId204Response struct {
}

func (response DeleteApiTasksId204Response) VisitDeleteApiTasksIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteApiTasksId404Response struct {
}

func (response DeleteApiTasksId404Response) VisitDeleteApiTasksIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PatchApiTasksIdRequestObject struct {
	Id   uint `json:"id"`
	Body *PatchApiTasksIdJSONRequestBody
}

type PatchApiTasksIdResponseObject interface {
	VisitPatchApiTasksIdResponse(w http.ResponseWriter) error
}

type PatchApiTasksId200JSONResponse Message

func (response PatchApiTasksId200JSONResponse) VisitPatchApiTasksIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PatchApiTasksId404Response struct {
}

func (response PatchApiTasksId404Response) VisitPatchApiTasksIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get all tasks
	// (GET /api/tasks)
	GetApiTasks(ctx context.Context, request GetApiTasksRequestObject) (GetApiTasksResponseObject, error)
	// Create a new task
	// (POST /api/tasks)
	PostApiTasks(ctx context.Context, request PostApiTasksRequestObject) (PostApiTasksResponseObject, error)
	// Delete a task by ID
	// (DELETE /api/tasks/{id})
	DeleteApiTasksId(ctx context.Context, request DeleteApiTasksIdRequestObject) (DeleteApiTasksIdResponseObject, error)
	// Update a task by ID
	// (PATCH /api/tasks/{id})
	PatchApiTasksId(ctx context.Context, request PatchApiTasksIdRequestObject) (PatchApiTasksIdResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetApiTasks operation middleware
func (sh *strictHandler) GetApiTasks(ctx echo.Context) error {
	var request GetApiTasksRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetApiTasks(ctx.Request().Context(), request.(GetApiTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetApiTasks")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetApiTasksResponseObject); ok {
		return validResponse.VisitGetApiTasksResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostApiTasks operation middleware
func (sh *strictHandler) PostApiTasks(ctx echo.Context) error {
	var request PostApiTasksRequestObject

	var body PostApiTasksJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostApiTasks(ctx.Request().Context(), request.(PostApiTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostApiTasks")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostApiTasksResponseObject); ok {
		return validResponse.VisitPostApiTasksResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteApiTasksId operation middleware
func (sh *strictHandler) DeleteApiTasksId(ctx echo.Context, id uint) error {
	var request DeleteApiTasksIdRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteApiTasksId(ctx.Request().Context(), request.(DeleteApiTasksIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteApiTasksId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteApiTasksIdResponseObject); ok {
		return validResponse.VisitDeleteApiTasksIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PatchApiTasksId operation middleware
func (sh *strictHandler) PatchApiTasksId(ctx echo.Context, id uint) error {
	var request PatchApiTasksIdRequestObject

	request.Id = id

	var body PatchApiTasksIdJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchApiTasksId(ctx.Request().Context(), request.(PatchApiTasksIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchApiTasksId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchApiTasksIdResponseObject); ok {
		return validResponse.VisitPatchApiTasksIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

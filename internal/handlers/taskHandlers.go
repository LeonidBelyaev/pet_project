package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"pet_project/internal/taskService"
	"pet_project/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

func (h *Handler) DeleteApiTasksId(ctx context.Context, request tasks.DeleteApiTasksIdRequestObject) (tasks.DeleteApiTasksIdResponseObject, error) {
	err := h.Service.DeleteTaskById(request.Id)
	if err != nil {
		return nil, err
	}
	return tasks.DeleteApiTasksId204Response{}, nil
}

func (h *Handler) PatchApiTasksId(ctx context.Context, request tasks.PatchApiTasksIdRequestObject) (tasks.PatchApiTasksIdResponseObject, error) {
	taskRequest := request.Body
	if taskRequest == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	taskToUpdate := taskService.Message{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	updatedTask, err := h.Service.UpdateTaskById(request.Id, taskToUpdate)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "task not found")
	}

	response := tasks.PatchApiTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}
	return response, nil
}

func (h *Handler) GetApiTasks(_ context.Context, _ tasks.GetApiTasksRequestObject) (tasks.GetApiTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTask()
	if err != nil {
		return nil, err
	}

	response := tasks.GetApiTasks200JSONResponse{}
	for _, tsk := range allTasks {
		task := tasks.Message{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *Handler) PostApiTasks(_ context.Context, request tasks.PostApiTasksRequestObject) (tasks.PostApiTasksResponseObject, error) {
	taskRequest := request.Body
	taskToCreate := taskService.Message{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	response := tasks.PostApiTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	return response, nil
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{Service: service}
}

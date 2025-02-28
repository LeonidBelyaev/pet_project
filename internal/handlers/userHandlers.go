package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"pet_project/internal/taskService"
	"pet_project/internal/userService"
	"pet_project/internal/web/users"
)

type UsersHandler struct {
	UsersService *userService.UsersService
	TaskService  *taskService.TaskService
}

func (u UsersHandler) GetUsersIdTasks(ctx context.Context, request users.GetUsersIdTasksRequestObject) (users.GetUsersIdTasksResponseObject, error) {
	tasks, err := u.UsersService.GetTasksForUser(request.Id)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := users.GetUsersIdTasks200JSONResponse{}
	for _, task := range tasks {
		taskResponse := users.Message{
			Id:     &task.ID,
			Task:   task.Task,
			IsDone: &task.IsDone,
			UserId: task.UserID,
		}
		response = append(response, taskResponse)
	}
	return response, nil
}

func (u UsersHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.UsersService.GetAllUsers()
	if err != nil {
		return nil, err
	}
	response := users.GetUsers200JSONResponse{}
	for _, usr := range allUsers {
		user := users.Users{
			Id:       &usr.ID,
			Email:    usr.Email,
			Password: usr.Password,
		}
		response = append(response, user)
	}
	return response, nil
}

func (u UsersHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := userService.Users{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
	createdUser, err := u.UsersService.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}
	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    createdUser.Email,
		Password: createdUser.Password,
	}
	return response, nil
}

func (u UsersHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	err := u.UsersService.DeleteUserById(request.Id)
	if err != nil {
		return nil, err
	}
	return users.DeleteUsersId204Response{}, nil
}

func (u UsersHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userRequest := request.Body
	if userRequest == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	userToUpdate := userService.Users{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	updatedUser, err := u.UsersService.UpdateUserById(request.Id, userToUpdate)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	response := users.PatchUsersId200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    updatedUser.Email,
		Password: updatedUser.Password,
	}
	return response, nil
}

func NewUsersHandlers(userService *userService.UsersService, taskService *taskService.TaskService) *UsersHandler {
	return &UsersHandler{
		UsersService: userService,
		TaskService:  taskService,
	}
}

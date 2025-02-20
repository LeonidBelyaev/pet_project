package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"pet_project/internal/database"
	"pet_project/internal/handlers"
	"pet_project/internal/taskService"
	"pet_project/internal/userService"
	"pet_project/internal/web/tasks"
	"pet_project/internal/web/users"
)

func main() {
	database.InitDB()
	//database.DB.AutoMigrate(&taskService.Message{})

	messageRepo := taskService.NewTaskRepository(database.DB)
	messageService := taskService.NewService(messageRepo)

	messageHandler := handlers.NewHandler(messageService)

	usersRepo := userService.NewUsersRepository(database.DB)
	usersService := userService.NewUsersService(usersRepo)

	usersHandler := handlers.NewUsersHandlers(usersService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictUsersHandler := users.NewStrictHandler(usersHandler, nil)
	users.RegisterHandlers(e, strictUsersHandler)

	strictHandler := tasks.NewStrictHandler(messageHandler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}

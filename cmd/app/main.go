package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"pet_project/internal/database"
	"pet_project/internal/handlers"
	"pet_project/internal/taskService"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Message{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/tasks", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", handler.UpdateTaskById).Methods("PATCH")
	router.HandleFunc("/api/tasks/{id}", handler.DeleteTaskById).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}

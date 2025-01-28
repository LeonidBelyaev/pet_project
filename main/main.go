package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var task string

type requestBody struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var task []Message
		DB.Find(&task)
		json.NewEncoder(w).Encode(task)
	} else {
		fmt.Fprintln(w, "Поддерживается только метод GET")
	}
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	var req requestBody
	if r.Method == http.MethodPost {
		json.NewDecoder(r.Body).Decode(&req)
		task = req.Message
		message := Message{Task: task, IsDone: false}
		DB.Create(&message)
		json.NewEncoder(w).Encode(message)
	}
}

func main() {
	InitDB()
	DB.AutoMigrate(&Message{})
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", TaskHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}

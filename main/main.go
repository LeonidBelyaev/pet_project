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

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {
		var req requestBody
		vars := mux.Vars(r)
		id := vars["id"]
		json.NewDecoder(r.Body).Decode(&req)

		var message Message
		DB.First(&message, id)
		if req.Message != "" {
			message.Task = req.Message
		}
		if req.Message != "" {
			message.IsDone = true
		}
		DB.Save(&message)
		json.NewEncoder(w).Encode(message)
	}
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		var req requestBody
		vars := mux.Vars(r)
		id := vars["id"]
		json.NewDecoder(r.Body).Decode(&req)
		var message Message
		DB.Delete(&message, id)
	}
}

func main() {
	InitDB()
	DB.AutoMigrate(&Message{})
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", TaskHandler).Methods("POST")
	router.HandleFunc("/api/task/{id}", PatchHandler).Methods("PATCH")
	router.HandleFunc("/api/task/{id}", DeleteHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}

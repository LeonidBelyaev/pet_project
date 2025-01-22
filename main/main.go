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
		fmt.Fprintln(w, "Hello, ", task)
	} else {
		fmt.Fprintln(w, "Поддерживается только метод GET")
	}
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	var req requestBody
	if r.Method == http.MethodPost {
		json.NewDecoder(r.Body).Decode(&req)
		task = req.Message
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", TaskHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}

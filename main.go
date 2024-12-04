package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var task Task

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if task.Title != "" {
		fmt.Fprintf(w, "hello, %s ", task.Title)
	} else {
		fmt.Fprintf(w, "task does not has title")
	}

}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	requestBody := r.Body
	err := json.NewDecoder(requestBody).Decode(&task)

	if err != nil {
		fmt.Println(err)
	}

	taskJSON, _ := json.Marshal(task)

	fmt.Fprintf(w, "Task %v?", task)
	fmt.Fprintf(w, "Task %s in json format", taskJSON)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", TaskHandler).Methods("POST")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		fmt.Println(err)
	}
}

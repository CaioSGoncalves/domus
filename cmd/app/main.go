package main

import (
	"fmt"
	"net/http"

	"home-todo/internal/handlers"
	"home-todo/internal/repository"
	"home-todo/internal/services"

	"github.com/a-h/templ"
)

func main() {
	repo := repository.NewTaskRepo()
	service := services.NewTaskService(repo)
	handler := handlers.NewTaskHandler(service)

	http.Handle("/", templ.Handler(handler.Home()))
	http.HandleFunc("/tasks", handler.AddTask)
	http.HandleFunc("/tasks/move", handler.MoveTask)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

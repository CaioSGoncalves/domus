package main

import (
	"fmt"
	"net/http"

	"domus/internal/handler"
	"domus/internal/repository"
	"domus/internal/service"
)

func main() {
	http.Handle("/", handler.SPAHandler())

	taskRepo := repository.NewMockTaskRepo()
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)
	http.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			taskHandler.GetTasks(w, r)
			return
		}

		if r.Method == http.MethodPost {
			taskHandler.AddTask(w, r)
			return
		}

	})
	http.HandleFunc("/api/tasks/move", taskHandler.MoveTask)

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

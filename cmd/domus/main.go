package main

import (
	"net/http"

	"domus/internal/server"
)

func main() {
	http.Handle("/", server.SPAHandler())

	http.HandleFunc("/api/tasks", server.GetTasks)
	http.HandleFunc("/api/tasks/move", server.MoveTask)

	http.ListenAndServe(":8080", nil)
}

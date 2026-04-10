package handlers

import (
	"net/http"
	"strconv"

	"home-todo/internal/services"
	"home-todo/internal/views"

	"github.com/a-h/templ"
)

type TaskHandler struct {
	service *services.TaskService
}

func NewTaskHandler(s *services.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) Home() templ.Component {
	tasks := h.service.ListTasks()
	return views.Page(tasks)
}

func (h *TaskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
	task := r.FormValue("task")

	t := h.service.AddTask(task)

	views.TaskItem(t).Render(r.Context(), w)
}

func (h *TaskHandler) MoveTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	status := r.FormValue("status")

	id, _ := strconv.Atoi(idStr)

	h.service.MoveTask(id, status)

	// re-renderiza TODAS as colunas
	tasks := h.service.ListTasks()
	views.Columns(tasks).Render(r.Context(), w)
}

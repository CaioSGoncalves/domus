package handler

import (
	"domus/internal/service"
	"encoding/json"
	"net/http"
)

type TaskHandler struct {
	svc *service.TaskService
}

func NewTaskHandler(svc *service.TaskService) *TaskHandler {
	return &TaskHandler{svc}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.svc.GetTasks()
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) MoveTask(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

type AddTaskRequest struct {
	PanelID  int    `json:"panelId"`
	ColumnID int    `json:"columnId"`
	Title    string `json:"title"`
}

func (h *TaskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
	var req AddTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		http.Error(w, "title required", http.StatusBadRequest)
		return
	}

	task, err := h.svc.AddTask(req.PanelID, req.ColumnID, req.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(task)
}

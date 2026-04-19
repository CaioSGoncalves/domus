package service

import (
	"domus/internal/domain"
)

type Repository interface {
	GetTasks() []domain.Panel
	AddTask(pID int, cID int, title string) (*domain.Task, error)
}

type TaskService struct {
	repo Repository
}

func NewTaskService(repo Repository) *TaskService {
	return &TaskService{repo}
}

func (s *TaskService) GetTasks() []domain.Panel {
	return s.repo.GetTasks()
}

func (s *TaskService) AddTask(panelID int, columnID int, title string) (*domain.Task, error) {
	task, err := s.repo.AddTask(panelID, columnID, title)

	return task, err

}

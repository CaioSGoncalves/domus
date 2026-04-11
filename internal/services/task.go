package services

import "domus/internal/repository"

type TaskService struct {
	repo *repository.TaskRepo
}

func NewTaskService(r *repository.TaskRepo) *TaskService {
	return &TaskService{repo: r}
}

func (s *TaskService) AddTask(name string) repository.Task {
	return s.repo.Add(name)
}

func (s *TaskService) MoveTask(id int, status string) {
	s.repo.Move(id, status)
}

func (s *TaskService) ListTasks() []repository.Task {
	return s.repo.List()
}

package repository

import (
	"domus/internal/domain"
	"errors"
)

var data = []domain.Panel{
	{
		ID:   1,
		Name: "Casa",
		Columns: []domain.Column{
			{ID: 1, Title: "A Fazer", Tasks: []domain.Task{{ID: 1, Title: "Lavar louça"}}},
			{ID: 2, Title: "Em Progresso", Tasks: []domain.Task{}},
			{ID: 3, Title: "Feito", Tasks: []domain.Task{}},
		},
	},
	{
		ID:   2,
		Name: "Homelab",
		Columns: []domain.Column{
			{ID: 4, Title: "A Fazer", Tasks: []domain.Task{{ID: 2, Title: "Domus: criar Tarefas"}}},
			{ID: 5, Title: "Em Progresso", Tasks: []domain.Task{}},
			{ID: 6, Title: "Feito", Tasks: []domain.Task{}},
		},
	},
	{
		ID:   3,
		Name: "Supermercado",
		Columns: []domain.Column{
			{ID: 7, Title: "Lista Padrão", Tasks: []domain.Task{{ID: 3, Title: "Mussarela"}}},
			{ID: 8, Title: "Para Comprar", Tasks: []domain.Task{}},
			{ID: 9, Title: "Na Cesta", Tasks: []domain.Task{}},
		},
	},
}

var nextTaskID int = 4

type MockTaskRepo struct {
	data []domain.Panel
}

func NewMockTaskRepo() *MockTaskRepo {
	return &MockTaskRepo{data: data}
}

func (r *MockTaskRepo) GetTasks() []domain.Panel {
	return r.data
}

func (r *MockTaskRepo) AddTask(pID int, cID int, title string) (*domain.Task, error) {
	for iP, p := range r.data {
		if p.ID == pID {
			for iC, c := range p.Columns {
				if c.ID == cID {
					t := domain.Task{
						ID:    nextTaskID,
						Title: title,
					}
					r.data[iP].Columns[iC].Tasks = append(c.Tasks, t)
					return &t, nil
				}
			}
		}
	}
	return nil, errors.New("Invalid Panel or Column ID")
}

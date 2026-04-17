package server

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type Column struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Tasks []Task `json:"tasks"`
}

type Panel struct {
	Name    string   `json:"name"`
	Columns []Column `json:"columns"`
}

var data = []Panel{
	{
		Name: "Casa",
		Columns: []Column{
			{ID: "todo", Title: "A Fazer", Tasks: []Task{{ID: "1", Title: "Lavar louça"}}},
			{ID: "doing", Title: "Em Progresso", Tasks: []Task{}},
			{ID: "done", Title: "Feito", Tasks: []Task{}},
		},
	},
	{
		Name: "Homelab",
		Columns: []Column{
			{ID: "todo", Title: "A Fazer", Tasks: []Task{{ID: "1", Title: "Domus: criar Tarefas"}}},
			{ID: "doing", Title: "Em Progresso", Tasks: []Task{}},
			{ID: "done", Title: "Feito", Tasks: []Task{}},
		},
	},
	{
		Name: "Supermercado",
		Columns: []Column{
			{ID: "todo", Title: "Lista Padrão", Tasks: []Task{{ID: "1", Title: "Papel Higiênico"}}},
			{ID: "doing", Title: "Para Comprar", Tasks: []Task{}},
			{ID: "done", Title: "Na Cesta", Tasks: []Task{}},
		},
	},
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(data)
}

func MoveTask(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

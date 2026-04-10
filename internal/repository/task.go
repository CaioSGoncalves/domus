package repository

type Task struct {
	ID     int
	Name   string
	Status string
}

type TaskRepo struct {
	tasks  []Task
	nextID int
}

func NewTaskRepo() *TaskRepo {
	return &TaskRepo{
		tasks: []Task{
			{ID: 1, Name: "Teste 1", Status: "todo"},
			{ID: 2, Name: "Teste 2", Status: "doing"},
		},
		nextID: 3,
	}
}

func (r *TaskRepo) Add(name string) Task {
	task := Task{
		ID:     r.nextID,
		Name:   name,
		Status: "todo",
	}

	r.nextID++
	r.tasks = append(r.tasks, task)

	return task
}

func (r *TaskRepo) List() []Task {
	return r.tasks
}

func (r *TaskRepo) Move(id int, status string) {
	for i := range r.tasks {
		if r.tasks[i].ID == id {
			r.tasks[i].Status = status
			return
		}
	}
}

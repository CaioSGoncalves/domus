package domain

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Column struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Tasks []Task `json:"tasks"`
}

type Panel struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Columns []Column `json:"columns"`
}

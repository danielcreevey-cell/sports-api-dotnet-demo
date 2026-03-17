package dto

// TodoListDto is a representation of a todo list with its items.
type TodoListDto struct {
	ID     int           `json:"id"`
	Title  string        `json:"title"`
	Colour string        `json:"colour"`
	Items  []TodoItemDto `json:"items"`
}

// TodosVm is the view model for the todos endpoint.
type TodosVm struct {
	PriorityLevels []LookupDto   `json:"priorityLevels"`
	Lists          []TodoListDto `json:"lists"`
}

package domain

// TodoList represents a collection of todo items.
type TodoList struct {
	BaseAuditableEntity
	Title  string
	Colour Colour
	Items  []TodoItem
}

// NewTodoList creates a new TodoList with default values.
func NewTodoList() TodoList {
	return TodoList{
		Colour: ColourWhite,
	}
}

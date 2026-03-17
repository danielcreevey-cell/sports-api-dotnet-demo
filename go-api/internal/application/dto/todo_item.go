package dto

// TodoItemBriefDto is a brief representation of a todo item.
type TodoItemBriefDto struct {
	ID     int    `json:"id"`
	ListID int    `json:"listId"`
	Title  string `json:"title"`
	Done   bool   `json:"done"`
}

// TodoItemDto is a detailed representation of a todo item.
type TodoItemDto struct {
	ID       int    `json:"id"`
	ListID   int    `json:"listId"`
	Title    string `json:"title"`
	Done     bool   `json:"done"`
	Priority int    `json:"priority"`
	Note     string `json:"note"`
}

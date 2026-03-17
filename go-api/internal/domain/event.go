package domain

// DomainEvent is the interface all domain events implement.
type DomainEvent interface {
	EventName() string
}

// TodoItemCreatedEvent is raised when a new TodoItem is created.
type TodoItemCreatedEvent struct {
	Item *TodoItem
}

// EventName returns the name of the event.
func (e TodoItemCreatedEvent) EventName() string { return "TodoItemCreatedEvent" }

// TodoItemCompletedEvent is raised when a TodoItem transitions to done.
type TodoItemCompletedEvent struct {
	Item *TodoItem
}

// EventName returns the name of the event.
func (e TodoItemCompletedEvent) EventName() string { return "TodoItemCompletedEvent" }

// TodoItemDeletedEvent is raised when a TodoItem is deleted.
type TodoItemDeletedEvent struct {
	Item *TodoItem
}

// EventName returns the name of the event.
func (e TodoItemDeletedEvent) EventName() string { return "TodoItemDeletedEvent" }

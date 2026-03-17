package domain

import "time"

// TodoItem represents a single todo item within a list.
type TodoItem struct {
	BaseAuditableEntity
	ListID   int
	Title    string
	Note     string
	Priority PriorityLevel
	Reminder *time.Time
	done     bool
	List     *TodoList
}

// Done returns the current done status.
func (t *TodoItem) Done() bool { return t.done }

// SetDone sets the done status. If transitioning from false to true, raises TodoItemCompletedEvent.
func (t *TodoItem) SetDone(value bool) {
	if value && !t.done {
		t.AddDomainEvent(TodoItemCompletedEvent{Item: t})
	}
	t.done = value
}

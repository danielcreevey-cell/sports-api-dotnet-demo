package entities

import (
	"time"

	"github.com/danielcreevey-cell/sports-api-dotnet-demo/internal/domain/enums"
	"github.com/danielcreevey-cell/sports-api-dotnet-demo/internal/domain/events"
)

// TodoItem is a single entry in a TodoList.
//
// Ported from src/Domain/Entities/TodoItem.cs. In .NET the Done property
// has a custom setter that raises a TodoItemCompletedEvent on the
// not-done → done transition; in Go we achieve the same behaviour by
// exposing SetDone/Done methods over an unexported field.
type TodoItem struct {
	BaseAuditableEntity

	ListID   int
	Title    *string
	Note     *string
	Priority enums.PriorityLevel
	Reminder *time.Time

	done bool

	// List is the back-reference to the owning TodoList. It may be nil
	// when the item has not yet been attached to a list (mirrors the
	// .NET `null!` late-initialised navigation property).
	List *TodoList
}

// GetID satisfies the events.TodoItemRef interface so that *TodoItem
// can be embedded in domain events without causing an import cycle
// between the entities and events packages.
func (t *TodoItem) GetID() int { return t.ID }

// Done reports whether the item is marked complete.
func (t *TodoItem) Done() bool { return t.done }

// SetDone updates the done flag. When the value transitions from false
// to true, a TodoItemCompletedEvent is appended to the item's domain
// events queue.
func (t *TodoItem) SetDone(val bool) {
	if val && !t.done {
		t.AddDomainEvent(events.NewTodoItemCompletedEvent(t))
	}
	t.done = val
}

// Package events contains domain events for the CleanArchitecture domain layer.
//
// Ported from src/Domain/Common/BaseEvent.cs and src/Domain/Events/*.cs.
package events

// DomainEvent is the marker interface implemented by every domain event.
//
// Equivalent to the .NET BaseEvent abstract class, which derives from
// MediatR's INotification. In Go we use a marker interface with an
// unexported method so that only types in this package (or types that
// embed BaseDomainEvent) can satisfy it.
type DomainEvent interface {
	isDomainEvent()
}

// BaseDomainEvent can be embedded by concrete event structs to satisfy
// the DomainEvent interface.
type BaseDomainEvent struct{}

func (BaseDomainEvent) isDomainEvent() {}

// TodoItemRef is the minimal contract for a TodoItem referenced from a
// domain event. It exists to break what would otherwise be an import
// cycle between the events and entities packages: entities depends on
// events (to emit events), and events need to reference the TodoItem
// aggregate without importing entities.
//
// The *entities.TodoItem type satisfies this interface.
type TodoItemRef interface {
	GetID() int
}

// TodoItemCreatedEvent is raised when a new TodoItem has been created.
//
// Ported from src/Domain/Events/TodoItemCreatedEvent.cs.
type TodoItemCreatedEvent struct {
	BaseDomainEvent
	Item TodoItemRef
}

// NewTodoItemCreatedEvent constructs a TodoItemCreatedEvent.
func NewTodoItemCreatedEvent(item TodoItemRef) TodoItemCreatedEvent {
	return TodoItemCreatedEvent{Item: item}
}

// TodoItemCompletedEvent is raised when a TodoItem transitions from
// not-done to done.
//
// Ported from src/Domain/Events/TodoItemCompletedEvent.cs.
type TodoItemCompletedEvent struct {
	BaseDomainEvent
	Item TodoItemRef
}

// NewTodoItemCompletedEvent constructs a TodoItemCompletedEvent.
func NewTodoItemCompletedEvent(item TodoItemRef) TodoItemCompletedEvent {
	return TodoItemCompletedEvent{Item: item}
}

// TodoItemDeletedEvent is raised when a TodoItem has been deleted.
//
// Ported from src/Domain/Events/TodoItemDeletedEvent.cs.
type TodoItemDeletedEvent struct {
	BaseDomainEvent
	Item TodoItemRef
}

// NewTodoItemDeletedEvent constructs a TodoItemDeletedEvent.
func NewTodoItemDeletedEvent(item TodoItemRef) TodoItemDeletedEvent {
	return TodoItemDeletedEvent{Item: item}
}

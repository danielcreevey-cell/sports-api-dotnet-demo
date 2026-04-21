package entities_test

import (
	"testing"

	"github.com/danielcreevey-cell/sports-api-dotnet-demo/internal/domain/entities"
	"github.com/danielcreevey-cell/sports-api-dotnet-demo/internal/domain/events"
)

func TestTodoItem_SetDone_EmitsEventOnTransitionToDone(t *testing.T) {
	item := &entities.TodoItem{}

	item.SetDone(true)

	emitted := item.DomainEvents()
	if len(emitted) != 1 {
		t.Fatalf("len(DomainEvents) = %d, want 1", len(emitted))
	}
	completed, ok := emitted[0].(events.TodoItemCompletedEvent)
	if !ok {
		t.Fatalf("event type = %T, want TodoItemCompletedEvent", emitted[0])
	}
	if completed.Item != item {
		t.Errorf("event.Item = %v, want %v", completed.Item, item)
	}
	if !item.Done() {
		t.Error("Done() = false after SetDone(true), want true")
	}
}

func TestTodoItem_SetDone_DoesNotEmitEventWhenAlreadyDone(t *testing.T) {
	item := &entities.TodoItem{}

	item.SetDone(true)
	item.ClearDomainEvents()
	item.SetDone(true)

	if got := item.DomainEvents(); len(got) != 0 {
		t.Errorf("len(DomainEvents) after redundant SetDone(true) = %d, want 0", len(got))
	}
}

func TestTodoItem_SetDone_DoesNotEmitEventOnTransitionToNotDone(t *testing.T) {
	item := &entities.TodoItem{}

	item.SetDone(true)
	item.ClearDomainEvents()
	item.SetDone(false)

	if got := item.DomainEvents(); len(got) != 0 {
		t.Errorf("len(DomainEvents) after SetDone(false) = %d, want 0", len(got))
	}
	if item.Done() {
		t.Error("Done() = true after SetDone(false), want false")
	}
}

func TestTodoItem_SetDone_DoesNotEmitEventWhenStaysFalse(t *testing.T) {
	item := &entities.TodoItem{}

	item.SetDone(false)

	if got := item.DomainEvents(); len(got) != 0 {
		t.Errorf("len(DomainEvents) = %d, want 0", len(got))
	}
}

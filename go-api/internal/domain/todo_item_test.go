package domain

import "testing"

func TestSetDone_FalseToTrue_RaisesCompletedEvent(t *testing.T) {
	item := &TodoItem{}
	item.SetDone(true)

	events := item.DomainEvents()
	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}

	completed, ok := events[0].(TodoItemCompletedEvent)
	if !ok {
		t.Fatalf("expected TodoItemCompletedEvent, got %T", events[0])
	}
	if completed.Item != item {
		t.Error("event item does not reference the original item")
	}
}

func TestSetDone_False_DoesNotRaiseEvent(t *testing.T) {
	item := &TodoItem{}
	item.SetDone(false)

	events := item.DomainEvents()
	if len(events) != 0 {
		t.Fatalf("expected 0 events, got %d", len(events))
	}
}

func TestSetDone_TrueToTrue_DoesNotRaiseDuplicateEvent(t *testing.T) {
	item := &TodoItem{}
	item.SetDone(true)
	item.SetDone(true) // already done, should not raise again

	events := item.DomainEvents()
	if len(events) != 1 {
		t.Fatalf("expected 1 event (no duplicate), got %d", len(events))
	}
}

func TestDone_ReturnsFalseByDefault(t *testing.T) {
	item := &TodoItem{}
	if item.Done() {
		t.Error("expected Done() to be false by default")
	}
}

func TestSetDone_TrueToFalse_DoesNotRaiseEvent(t *testing.T) {
	item := &TodoItem{}
	item.SetDone(true)
	item.ClearDomainEvents()
	item.SetDone(false)

	events := item.DomainEvents()
	if len(events) != 0 {
		t.Fatalf("expected 0 events when setting done to false, got %d", len(events))
	}
}

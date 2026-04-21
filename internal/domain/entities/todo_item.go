package entities

import "time"

// PriorityLevel represents the priority of a TodoItem.
type PriorityLevel int

const (
	PriorityNone PriorityLevel = iota
	PriorityLow
	PriorityMedium
	PriorityHigh
)

// TodoItem is the Go port of the Domain TodoItem entity.
type TodoItem struct {
	ID       int
	ListID   int
	Title    *string
	Note     *string
	Priority PriorityLevel
	Reminder *time.Time
	Done     bool
}

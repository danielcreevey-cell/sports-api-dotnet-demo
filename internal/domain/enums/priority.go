// Package enums contains enumeration types for the CleanArchitecture
// domain layer.
package enums

// PriorityLevel indicates the priority of a TodoItem.
//
// Ported from src/Domain/Enums/PriorityLevel.cs.
type PriorityLevel int

const (
	PriorityNone   PriorityLevel = 0
	PriorityLow    PriorityLevel = 1
	PriorityMedium PriorityLevel = 2
	PriorityHigh   PriorityLevel = 3
)

// String returns a human-readable name for the priority level.
func (p PriorityLevel) String() string {
	switch p {
	case PriorityNone:
		return "None"
	case PriorityLow:
		return "Low"
	case PriorityMedium:
		return "Medium"
	case PriorityHigh:
		return "High"
	default:
		return "Unknown"
	}
}

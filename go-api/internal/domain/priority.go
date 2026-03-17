package domain

// PriorityLevel represents the priority of a todo item.
type PriorityLevel int

const (
	PriorityNone   PriorityLevel = 0
	PriorityLow    PriorityLevel = 1
	PriorityMedium PriorityLevel = 2
	PriorityHigh   PriorityLevel = 3
)

// String returns the string representation of the priority level.
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
		return "None"
	}
}

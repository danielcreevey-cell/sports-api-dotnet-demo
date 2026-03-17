package domain

import "testing"

func TestPriorityLevel_String(t *testing.T) {
	tests := []struct {
		level    PriorityLevel
		expected string
	}{
		{PriorityNone, "None"},
		{PriorityLow, "Low"},
		{PriorityMedium, "Medium"},
		{PriorityHigh, "High"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.level.String(); got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestPriorityLevel_String_Unknown(t *testing.T) {
	unknown := PriorityLevel(99)
	if got := unknown.String(); got != "None" {
		t.Errorf("expected unknown priority to return 'None', got %q", got)
	}
}

package domain

import "fmt"

// UnsupportedColourError is returned when an unsupported colour code is used.
type UnsupportedColourError struct {
	Code string
}

// Error returns the error message.
func (e *UnsupportedColourError) Error() string {
	return fmt.Sprintf("Colour %q is unsupported.", e.Code)
}

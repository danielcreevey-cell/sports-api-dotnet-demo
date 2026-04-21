package errors

import "fmt"

// ValidationFailure represents a single validation failure for a property.
//
// Mirrors FluentValidation.Results.ValidationFailure from the .NET project.
type ValidationFailure struct {
	PropertyName string
	ErrorMessage string
}

// ValidationError is returned when one or more validation failures occur.
//
// Ported from src/Application/Common/Exceptions/ValidationException.cs. The
// Errors map is keyed by property name, matching the JSON shape emitted by
// the .NET implementation.
type ValidationError struct {
	Errors map[string][]string
}

// NewValidationError returns a ValidationError with an empty Errors map.
func NewValidationError() *ValidationError {
	return &ValidationError{Errors: map[string][]string{}}
}

// NewValidationErrorFromFailures groups the provided failures by property
// name, matching the behaviour of the C# constructor that takes an
// IEnumerable<ValidationFailure>.
func NewValidationErrorFromFailures(failures []ValidationFailure) *ValidationError {
	errs := map[string][]string{}
	for _, f := range failures {
		errs[f.PropertyName] = append(errs[f.PropertyName], f.ErrorMessage)
	}
	return &ValidationError{Errors: errs}
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	return fmt.Sprintf("one or more validation failures have occurred: %d propert(ies) failed validation", len(e.Errors))
}

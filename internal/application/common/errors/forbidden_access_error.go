package errors

// ForbiddenAccessError indicates an operation was rejected because the caller
// is not authorized.
//
// Ported from src/Application/Common/Exceptions/ForbiddenAccessException.cs.
type ForbiddenAccessError struct{}

// NewForbiddenAccessError returns a new ForbiddenAccessError.
func NewForbiddenAccessError() *ForbiddenAccessError {
	return &ForbiddenAccessError{}
}

// Error implements the error interface.
func (e *ForbiddenAccessError) Error() string {
	return "forbidden access"
}

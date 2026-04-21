// Package errors contains sentinel error types for the CleanArchitecture
// domain layer.
package errors

import "fmt"

// UnsupportedColourError is returned when a hex code is not one of the
// supported Colour value-object codes.
//
// Ported from src/Domain/Exceptions/UnsupportedColourException.cs.
type UnsupportedColourError struct {
	Code string
}

// NewUnsupportedColourError constructs an UnsupportedColourError for
// the given code.
func NewUnsupportedColourError(code string) *UnsupportedColourError {
	return &UnsupportedColourError{Code: code}
}

func (e *UnsupportedColourError) Error() string {
	return fmt.Sprintf("Colour %q is unsupported.", e.Code)
}

// NotFoundError indicates that an entity with the given key was not
// found. It is the Go analogue of Ardalis.GuardClauses.NotFoundException
// used in the Application layer's guard clauses.
type NotFoundError struct {
	EntityName string
	Key        any
}

// NewNotFoundError constructs a NotFoundError.
func NewNotFoundError(entityName string, key any) *NotFoundError {
	return &NotFoundError{EntityName: entityName, Key: key}
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s with key %v was not found.", e.EntityName, e.Key)
}

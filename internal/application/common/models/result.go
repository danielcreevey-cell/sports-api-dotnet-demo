package models

// Result represents the outcome of an operation, optionally carrying errors.
//
// Ported from src/Application/Common/Models/Result.cs.
type Result struct {
	Succeeded bool
	Errors    []string
}

// Success returns a successful Result with no errors.
func Success() Result {
	return Result{Succeeded: true, Errors: []string{}}
}

// Failure returns a failed Result carrying the provided errors.
func Failure(errors []string) Result {
	if errors == nil {
		errors = []string{}
	}
	return Result{Succeeded: false, Errors: errors}
}

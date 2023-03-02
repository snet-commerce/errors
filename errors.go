package errors

import (
	"encoding/json"
	"fmt"
)

// Error is custom error
type Error struct {
	code    string
	message string
	errors  []violation
}

// New constructs Error
func New(code, message string, violations ...violation) *Error {
	if len(violations) == 0 {
		violations = make([]violation, 0)
	}

	return &Error{
		code:    code,
		message: message,
		errors:  violations,
	}
}

// Violations add new violations to error
func (e *Error) Violations(violations ...violation) {
	e.errors = append(e.errors, violations...)
}

// Error implements error interface
func (e *Error) Error() string {
	return fmt.Sprintf("error code %s: %s", e.code, e.message)
}

// MarshalJSON implements json.Marshaler interface
func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Errors  []violation `json:"errors"`
	}{
		Code:    e.code,
		Message: e.message,
		Errors:  e.errors,
	})
}

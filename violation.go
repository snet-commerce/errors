package errors

import "encoding/json"

// violation represents inner error
type violation struct {
	reason  string
	message string
}

// Violation construct new violation
func Violation(reason, message string) violation {
	return violation{
		reason:  reason,
		message: message,
	}
}

// MarshalJSON implements json.Marshaler interface
func (v violation) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Reason  string `json:"reason"`
		Message string `json:"message"`
	}{
		Reason:  v.reason,
		Message: v.message,
	})
}

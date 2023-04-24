package utils

import (
	"fmt"
)

// Error is an interface for any execution errors.
// It is used to return errors from the API.
type Error interface {
	error
	Code() int32
}

type apiError struct {
	code    int32
	message string
}

func (a *apiError) Error() string {
	return a.message
}

func (a *apiError) Code() int32 {
	return a.code
}

// New creates a new API error with a code and a message
func New(code int32, message string, args ...interface{}) Error {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	return &apiError{code: code, message: message}
}

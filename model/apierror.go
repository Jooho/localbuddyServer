package model

import "fmt"

// Error json model
type ApiError struct {
	Code    string
	Message string
}

// String() implement
func (e ApiError) String() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// NewApiError creates ApiError object
func NewApiError(code, message string) *ApiError {
	return &ApiError{code, message}
}
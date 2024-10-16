package errors

import (
	"fmt"
)

type InternalError struct {
	CustomError
}

func NewInternalError(message string) *InternalError {
	internalError := &InternalError{
		CustomError: CustomError{
			Code: "internal-error",
			error: fmt.Errorf(message),
		},
	}
	return internalError
}

func NewInternalErrorWithCause(message string, cause error) *InternalError {
	internalError := NewInternalError(message)
	internalError.error = cause
	return internalError
}
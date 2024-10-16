package errors

import (
	"fmt"
)

type CustomError struct {
	Code string
	error
}

func NewCustomError(code string, message string) *CustomError {
	customError := &CustomError{
		Code: code,
		error: fmt.Errorf(message),
	}
	return customError
}

func NewCustomErrorWithCause(code string, message string, cause error) *CustomError {
	customError := NewCustomError(code, message)
	customError.error = cause
	return customError
}

func (ce CustomError) Error() string {
	return fmt.Sprintf("%s: %s", ce.Code, ce.error.Error())
}

package errors

import "errors"

type DomainError struct {
	CustomError
}

func NewDomainError(code string, message string) *DomainError {
	domainError := &DomainError{
		CustomError: CustomError{
			Code:  code,
			error: errors.New(message),
		},
	}
	return domainError
}

func NewDomainErrorWithCause(code string, message string, cause error) *DomainError {
	domainError := NewDomainError(code, message)
	domainError.error = cause
	return domainError
}

func (ce *DomainError) Is(target error) bool { return target == ce }

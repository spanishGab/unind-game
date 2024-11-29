package creatures

import "spanishgab/unind/src/errors"

var InvalidAttributeValueError *errors.InternalError = errors.NewInternalError(
	"attribute values must be greater than zero",
)

var CreatureDiedError *errors.DomainError = errors.NewDomainError(
	"creature-died",
	"creature died",
)

package creatures

import "spanishgab/unind/src/errors"

var InvalidAttributeValueError *errors.InternalError = errors.NewInternalError(
	"attribute values must be greater than zero",
)

var MismatchedWeaponTypeError *errors.InternalError = errors.NewInternalError(
	"the weapon type does not match with the expected initial weapon type for this character",
)

var CreatureDiedError *errors.DomainError = errors.NewDomainError(
	"creature-died",
	"creature died",
)

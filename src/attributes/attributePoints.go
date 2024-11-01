package attributes

import (
	"fmt"
	"spanishgab/unind/src/errors"
)

const (
	ATTRIBUTE_NOT_FOUND      = "attribute %q does not exist"
	INVALID_INCREASING_VALUE = "increasing value must be less than zero"
	INVALID_DECREASING_VALUE = "decreasing value must be greater than zero"
)

type Attribute string

const (
	ATTACK       Attribute = "attack"
	DEFENSE      Attribute = "defense"
	HEALTH       Attribute = "healh"
	INTELLIGENCE Attribute = "intelligence"
)

type AttributePoints struct {
	points map[Attribute]float64
}

func (ap *AttributePoints) Get(attribute Attribute) (float64, *errors.InternalError) {
	value, ok := ap.points[attribute]
	if !ok {
		return 0.0, errors.NewInternalError(
			fmt.Sprintf(ATTRIBUTE_NOT_FOUND, attribute),
		)
	}
	return value, nil
}

func (ap *AttributePoints) Has(attribute Attribute) bool {
	_, ok := ap.points[attribute]
	return ok
}

func (ap *AttributePoints) Increase(attribute Attribute, points float64) *errors.InternalError {
	if !ap.Has(attribute) {
		return errors.NewInternalError(
			fmt.Sprintf(ATTRIBUTE_NOT_FOUND, attribute),
		)
	}
	if points < 0 {
		return errors.NewInternalError(INVALID_INCREASING_VALUE)
	}
	ap.points[attribute] += points
	return nil
}

func (ap *AttributePoints) Decrease(attribute Attribute, points float64) *errors.InternalError {
	if !ap.Has(attribute) {
		return errors.NewInternalError(
			fmt.Sprintf(ATTRIBUTE_NOT_FOUND, attribute),
		)
	}
	if points > 0 {
		return errors.NewInternalError(INVALID_DECREASING_VALUE)
	}
	ap.points[attribute] += points
	return nil
}

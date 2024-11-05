package attributes

import (
	"fmt"
	"spanishgab/unind/src/errors"
)

const (
	ATTRIBUTE_NOT_FOUND      = "attribute %q does not exist"
	INVALID_ATTRIBUTE_VALUE  = "attribute values must be greater than zero"
	INVALID_INCREASING_VALUE = "increasing value must be greater than zero"
	INVALID_DECREASING_VALUE = "decreasing value must be less than zero"
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

func NewBattlePoints(attackPoints, defensePoints float64) (*AttributePoints, *errors.InternalError) {
	if err := checkAttributeValues(attackPoints, defensePoints); err != nil {
		return nil, err
	}

	return &AttributePoints{
		points: map[Attribute]float64{
			ATTACK:  attackPoints,
			DEFENSE: defensePoints,
		},
	}, nil
}

func NewHealthPoints(healthPoints float64) (*AttributePoints, *errors.InternalError) {
	if err := checkAttributeValues(healthPoints); err != nil {
		return nil, err
	}
	return &AttributePoints{
		points: map[Attribute]float64{
			HEALTH: healthPoints,
		},
	}, nil
}

func NewIntelligencePoints(intelligencePoints float64) (*AttributePoints, *errors.InternalError) {
	if err := checkAttributeValues(intelligencePoints); err != nil {
		return nil, err
	}
	return &AttributePoints{
		points: map[Attribute]float64{
			INTELLIGENCE: intelligencePoints,
		},
	}, nil
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

func checkAttributeValues(value ...float64) *errors.InternalError {
	for _, attributePoint := range value {
		if attributePoint < 0 {
			return errors.NewInternalError(INVALID_ATTRIBUTE_VALUE)
		}
	}
	return nil
}

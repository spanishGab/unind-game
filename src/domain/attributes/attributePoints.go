package attributes

import (
	"fmt"
	"spanishgab/unind/src/errors"
	"spanishgab/unind/src/utils"
)

const (
	ATTRIBUTE_NOT_FOUND      string = "attribute %q does not exist"
	INVALID_ATTRIBUTE_VALUE  string = "attribute values must be greater than zero"
	INVALID_INCREASING_VALUE string = "increasing value must be greater than zero"
	INVALID_DECREASING_VALUE string = "decreasing value must be less than zero"
)

type Attribute string

const (
	ATTACK       Attribute = "attack"
	DEFENSE      Attribute = "defense"
	HEALTH       Attribute = "healh"
	INTELLIGENCE Attribute = "intelligence"
	STRENGTH     Attribute = "strength"
)

type AttributePoints struct {
	points map[Attribute]float64
}

func New(
	attackPoints float64,
	defensePoints float64,
	healthPoints float64,
	intelligencePoints float64,
	strengthPoints float64,
) (*AttributePoints, *errors.InternalError) {
	if !utils.AreAllPositive(
		attackPoints,
		defensePoints,
		attackPoints,
		defensePoints,
		healthPoints,
		intelligencePoints,
		strengthPoints,
	) {
		return nil, errors.NewInternalError(INVALID_ATTRIBUTE_VALUE)
	}

	return &AttributePoints{
		points: map[Attribute]float64{
			ATTACK:       attackPoints,
			DEFENSE:      defensePoints,
			HEALTH:       healthPoints,
			INTELLIGENCE: intelligencePoints,
			STRENGTH:     strengthPoints,
		},
	}, nil
}

func NewBattlePoints(attackPoints, defensePoints float64) (*AttributePoints, *errors.InternalError) {
	if !utils.AreAllPositive(attackPoints, defensePoints) {
		return nil, errors.NewInternalError(INVALID_ATTRIBUTE_VALUE)
	}

	return &AttributePoints{
		points: map[Attribute]float64{
			ATTACK:  attackPoints,
			DEFENSE: defensePoints,
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

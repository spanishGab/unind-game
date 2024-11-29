package creatures

import (
	"spanishgab/unind/src/errors"
	"spanishgab/unind/src/utils"
)

type CreatureAttributes struct {
	strength     float64
	health       float64
	intelligence float64
}

func NewCreatureAttributes(health float64, intelligence float64, strength float64) (*CreatureAttributes, *errors.InternalError) {
	if !utils.AreAllPositive(health, intelligence, strength) {
		return nil, InvalidAttributeValueError
	}

	return &CreatureAttributes{
		health:       health,
		intelligence: intelligence,
		strength:     strength,
	}, nil
}

func (ca *CreatureAttributes) GetStrength() float64 {
	return ca.strength
}

func (ca *CreatureAttributes) GetHealth() float64 {
	return ca.health
}

func (ca *CreatureAttributes) GetIntelligence() float64 {
	return ca.intelligence
}

func (ca *CreatureAttributes) SetStrength(points float64) *errors.InternalError {
	if !utils.AreAllPositive(points) {
		return InvalidAttributeValueError
	}
	ca.strength = points
	return nil
}

func (ca *CreatureAttributes) SetHealth(points float64) *errors.InternalError {
	if !utils.AreAllPositive(points) {
		ca.health = 0
		return InvalidAttributeValueError
	}
	ca.health = points
	return nil
}

func (ca *CreatureAttributes) SetIntelligence(points float64) *errors.InternalError {
	if !utils.AreAllPositive(points) {
		return InvalidAttributeValueError
	}
	ca.intelligence = points
	return nil
}

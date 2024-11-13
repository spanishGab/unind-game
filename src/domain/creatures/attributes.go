package creatures

import (
	"spanishgab/unind/src/errors"
	"spanishgab/unind/src/utils"
)

const (
	INVALID_ATTRIBUTE_VALUE string = "attribute values must be greater than zero"
)

type CreatureAttributes struct {
	strength     float64
	health       float64
	intelligence float64
}

func NewCreatureAttributes(health float64, intelligence float64, strength float64) (*CreatureAttributes, *errors.InternalError) {
	if !utils.AreAllPositive(health, intelligence, strength) {
		return nil, errors.NewInternalError(INVALID_ATTRIBUTE_VALUE)
	}

	return &CreatureAttributes{
		health:       health,
		intelligence: intelligence,
		strength:     strength,
	}, nil
}

func (cap *CreatureAttributes) GetStrength() float64 {
	return cap.strength
}

func (cap *CreatureAttributes) GetHealth() float64 {
	return cap.health
}

func (cap *CreatureAttributes) GetIntelligence() float64 {
	return cap.intelligence
}

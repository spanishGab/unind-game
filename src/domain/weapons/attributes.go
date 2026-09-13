package weapons

import (
	"spanishgab/unind/src/errors"
	"spanishgab/unind/src/utils"
)

const (
	InvalidAttributeValue string = "attribute values must be greater than zero"
)

type WeaponAttributes struct {
	attack  float64
	defense float64
}

func NewWeaponAttributes(attackPoints float64, defensePoints float64) (*WeaponAttributes, *errors.InternalError) {
	if !utils.AreAllPositive(attackPoints, defensePoints) {
		return nil, errors.NewInternalError(InvalidAttributeValue)
	}

	return &WeaponAttributes{
		attack:  attackPoints,
		defense: defensePoints,
	}, nil
}

func (bap *WeaponAttributes) Attack() float64 {
	return bap.attack
}

func (bap *WeaponAttributes) Defense() float64 {
	return bap.defense
}

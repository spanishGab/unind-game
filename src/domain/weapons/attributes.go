package weapons

import (
	"spanishgab/unind/src/errors"
	"spanishgab/unind/src/utils"
)

const (
	INVALID_ATTRIBUTE_VALUE string = "attribute values must be greater than zero"
)

type WeaponAttributes struct {
	attack  float64
	defense float64
}

func NewWeaponAttributes(attackPoints float64, defensePoints float64) (*WeaponAttributes, *errors.InternalError) {
	if !utils.AreAllPositive(attackPoints, defensePoints) {
		return nil, errors.NewInternalError(INVALID_ATTRIBUTE_VALUE)
	}

	return &WeaponAttributes{
		attack:  attackPoints,
		defense: defensePoints,
	}, nil
}

func (bap *WeaponAttributes) GetAttack() float64 {
	return bap.attack
}

func (bap *WeaponAttributes) GetDefense() float64 {
	return bap.defense
}

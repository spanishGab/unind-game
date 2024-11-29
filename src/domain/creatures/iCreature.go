package creatures

import (
	"spanishgab/unind/src/domain/potions"
	"spanishgab/unind/src/domain/weapons"
	"spanishgab/unind/src/errors"
)

type ICreature interface {
	Attack() float64
	Defend() *errors.DomainError
	Heal(potion potions.Potion)
	Strengthen(potion potions.Potion)
	EquipLeftHand(weapon weapons.Weapon)
	EquipRightHand(weapon weapons.Weapon)
}

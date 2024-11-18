package creatures

import (
	"spanishgab/unind/src/domain/potions"
	"spanishgab/unind/src/domain/weapons"
)

type ICreature interface {
	Attack() float64
	Defend()
	Heal(potion potions.Potion)
	Strengthen(potion potions.Potion)
	EquipLeftHand(weapon weapons.Weapon)
	EquipRightHand(weapon weapons.Weapon)
}

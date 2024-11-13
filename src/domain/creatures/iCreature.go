package creatures

import (
	"spanishgab/unind/src/domain/potions"
	"spanishgab/unind/src/domain/weapons"
)

type ICreature interface {
	attack() float64
	defend() float64
	heal(potion potions.Potion)
	strengthen(potion potions.Potion)
	equipLeftHand(weapon weapons.Weapon)
	equipRightHand(weapon weapons.Weapon)
}

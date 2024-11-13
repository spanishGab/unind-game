package creatures

import (
	"spanishgab/unind/src/domain/weapons"
)

type Race string

const (
	HUMAN       Race = "human"
	WIZARD      Race = "wizard"
	ELF         Race = "elf"
	DWARF       Race = "dwarf"
	ORC         Race = "orc"
	SORCERER    Race = "sorcerer"
	SHADOW_LORD Race = "shadow-lord"
)

type Creature struct {
	name       string
	race       Race
	leftHand   *weapons.Weapon
	rightHand  *weapons.Weapon
	attributes *CreatureAttributes
}

func New(
	name string,
	race Race,
	leftHand *weapons.Weapon,
	rightHand *weapons.Weapon,
	attributes CreatureAttributes,
) *Creature {
	return &Creature{
		name:       name,
		race:       race,
		leftHand:   leftHand,
		rightHand:  rightHand,
		attributes: &attributes,
	}
}

func (c *Creature) Attack() float64 {
	weaponsStrength := getWeaponAttackStrength(c.leftHand) + getWeaponAttackStrength(c.rightHand)

	return (c.attributes.GetStrength() + weaponsStrength) * (c.attributes.GetIntelligence() / 100)
}

func getWeaponAttackStrength(w *weapons.Weapon) float64 {
	if w == nil {
		return 0
	}
	return w.GetBattlePoints().GetAttack()
}

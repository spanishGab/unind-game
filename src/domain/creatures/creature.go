package creatures

import (
	"math"
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
	weaponsStrength := weapons.GetAttackPower(c.leftHand) + weapons.GetAttackPower(c.rightHand)

	return c.getBattlePoints(weaponsStrength)
}

func (c *Creature) Defend() float64 {
	weaponsStrength := weapons.GetDefensePower(c.leftHand) + weapons.GetDefensePower(c.rightHand)

	return c.getBattlePoints(weaponsStrength)
}

func (c *Creature) getBattlePoints(weaponsStrength float64) float64 {
	return math.Ceil(
		(c.attributes.GetStrength() + weaponsStrength) * (c.attributes.GetIntelligence() / 100),
	)
}

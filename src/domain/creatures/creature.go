package creatures

import (
	"math"
	"spanishgab/unind/src/domain/potions"
	"spanishgab/unind/src/domain/weapons"
)

const (
	MAX_HEALTH_POINTS float64 = 100
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
	// TODO: re-implement Defend
	weaponsStrength := weapons.GetDefensePower(c.leftHand) + weapons.GetDefensePower(c.rightHand)

	return c.getBattlePoints(weaponsStrength)
}

func (c *Creature) Heal(potion *potions.Potion) {
	if newHealth := c.attributes.health + float64(potion.GetUpgradePoints()); newHealth > MAX_HEALTH_POINTS {
		c.attributes.health = MAX_HEALTH_POINTS
	} else {
		c.attributes.health = newHealth
	}
}

func (c *Creature) getBattlePoints(weaponsStrength float64) float64 {
	return math.Ceil(
		(c.attributes.GetStrength() + weaponsStrength) * (c.attributes.GetIntelligence() / 100),
	)
}

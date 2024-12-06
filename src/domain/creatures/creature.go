package creatures

import (
	"math"
	"spanishgab/unind/src/domain/potions"
	"spanishgab/unind/src/domain/weapons"
	"spanishgab/unind/src/errors"
)

const (
	MaxHealthPoints   float64 = 100
	MaxStrengthPoints float64 = 100
)

type Race string

const (
	Human      Race = "human"
	Wizard     Race = "wizard"
	Elf        Race = "elf"
	Dwarf      Race = "dwarf"
	Orc        Race = "orc"
	Sorcerer   Race = "sorcerer"
	ShadowLord Race = "shadow-lord"
)

type Creature struct {
	name       string
	race       Race
	leftHand   *weapons.Weapon
	rightHand  *weapons.Weapon
	attributes *CreatureAttributes
}

func new(
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

func NewHuman(name string)

func (c *Creature) Attack() float64 {
	weaponsStrength := weapons.GetAttackPower(c.leftHand) + weapons.GetAttackPower(c.rightHand)

	return c.getBattlePoints(weaponsStrength)
}

func (c *Creature) Defend(attack float64) *errors.DomainError {
	weaponsStrength := weapons.GetDefensePower(c.leftHand) + weapons.GetDefensePower(c.rightHand)

	if defense := c.getBattlePoints(weaponsStrength) - attack; defense < 0 {
		err := c.attributes.SetHealth(c.attributes.GetHealth() + defense)
		if err != nil {
			return CreatureDiedError
		}
	}
	return nil
}

func (c *Creature) Heal(potion *potions.Potion) {
	if newHealth := c.attributes.GetHealth() + float64(potion.GetUpgradePoints()); newHealth > MaxHealthPoints {
		c.attributes.SetHealth(MaxHealthPoints)
	} else {
		c.attributes.SetHealth(newHealth)
	}
}

func (c *Creature) Strengthen(potion *potions.Potion) {
	if newStrength := c.attributes.GetStrength() + float64(potion.GetUpgradePoints()); newStrength > MaxStrengthPoints {
		c.attributes.SetStrength(MaxStrengthPoints)
	} else {
		c.attributes.SetStrength(newStrength)
	}
}

func (c *Creature) EquipRightHand(weapon *weapons.Weapon) {
	c.rightHand = weapon
}

func (c *Creature) EquipLeftHand(weapon *weapons.Weapon) {
	c.leftHand = weapon
}

func (c *Creature) getBattlePoints(weaponsStrength float64) float64 {
	return math.Ceil(
		(c.attributes.GetStrength() + weaponsStrength) * (c.attributes.GetIntelligence() / 100),
	)
}

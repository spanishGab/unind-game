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

type ICreature interface {
	Attack() float64
	Defend(attack float64) *errors.DomainError
	Heal(potion *potions.Potion)
	Strengthen(potion *potions.Potion)
	EquipLeftHand(weapon *weapons.Weapon)
	EquipRightHand(weapon *weapons.Weapon)
}

type Creature struct {
	name       string
	race       Race
	leftHand   *weapons.Weapon
	rightHand  *weapons.Weapon
	attributes *CreatureAttributes
}

func NewCreature(
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
	weaponsStrength := weapons.AttackPower(c.leftHand) + weapons.AttackPower(c.rightHand)

	return c.getBattlePoints(weaponsStrength)
}

func (c *Creature) Defend(attack float64) *errors.DomainError {
	weaponsStrength := weapons.DefensePower(c.leftHand) + weapons.DefensePower(c.rightHand)

	if defense := c.getBattlePoints(weaponsStrength) - attack; defense < 0 {
		err := c.attributes.SetHealth(c.attributes.Health() + defense)
		if err != nil {
			return CreatureDiedError
		}
	}
	return nil
}

func (c *Creature) Heal(potion *potions.Potion) {
	if newHealth := c.attributes.Health() + float64(potion.UpgradePoints()); newHealth > MaxHealthPoints {
		c.attributes.SetHealth(MaxHealthPoints)
	} else {
		c.attributes.SetHealth(newHealth)
	}
}

func (c *Creature) Strengthen(potion *potions.Potion) {
	if newStrength := c.attributes.Strength() + float64(potion.UpgradePoints()); newStrength > MaxStrengthPoints {
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
		(c.attributes.Strength() + weaponsStrength) * (c.attributes.Intelligence() / 100),
	)
}

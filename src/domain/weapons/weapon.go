package weapons

import (
	"spanishgab/unind/src/errors"
)

type WeaponType string

const (
	Sword       WeaponType = "sword"
	Shield      WeaponType = "shield"
	BowAndArrow WeaponType = "bow and arrow"
	Axe         WeaponType = "axe"
	Staff       WeaponType = "staff"
	Wand        WeaponType = "wand"
	Dagger      WeaponType = "dagger"
)

type Weapon struct {
	name       string
	type_      WeaponType
	attributes *WeaponAttributes
}

func NewWeapon(
	name string,
	type_ WeaponType,
	attackPoints float64,
	defensePoints float64,
) (*Weapon, *errors.InternalError) {
	battlePoints, err := NewWeaponAttributes(attackPoints, defensePoints)
	if err != nil {
		return nil, err
	}

	return &Weapon{
		name:       name,
		type_:      type_,
		attributes: battlePoints,
	}, nil
}

func (w *Weapon) Name() string {
	return w.name
}

func (w *Weapon) Type() WeaponType {
	return w.type_
}

func (w *Weapon) Attributes() *WeaponAttributes {
	return w.attributes
}

func AttackPower(w *Weapon) float64 {
	if w == nil {
		return 0
	}
	return w.Attributes().Attack()
}

func DefensePower(w *Weapon) float64 {
	if w == nil {
		return 0
	}
	return w.Attributes().Defense()
}

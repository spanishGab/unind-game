package weapons

import (
	"spanishgab/unind/src/errors"
)

type WeaponType string

const (
	SWORD         WeaponType = "sword"
	SHIELD        WeaponType = "shield"
	BOW_AND_ARROW WeaponType = "bow and arrow"
	AXE           WeaponType = "axe"
	STAFF         WeaponType = "staff"
	WAND          WeaponType = "wand"
	DAGGER        WeaponType = "dagger"
)

type Weapon struct {
	name       string
	type_      WeaponType
	attributes *WeaponAttributes
}

func new(
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

func NewSword(name string, attackPoints float64, defensePoints float64) (*Weapon, *errors.InternalError) {
	return new(
		name,
		SWORD,
		attackPoints,
		defensePoints,
	)
}

func NewShield(name string, defensePoints float64) (*Weapon, *errors.InternalError) {
	return new(
		name,
		SHIELD,
		0,
		defensePoints,
	)
}

func NewBowAndArrow(name string, attackPoints float64) (*Weapon, *errors.InternalError) {
	return new(
		name,
		BOW_AND_ARROW,
		attackPoints,
		0,
	)
}

func NewAxe(name string, attackPoints float64) (*Weapon, *errors.InternalError) {
	return new(
		name,
		AXE,
		attackPoints,
		0,
	)
}

func NewStaff(name string, attackPoints float64, defensePoints float64) (*Weapon, *errors.InternalError) {
	return new(
		name,
		STAFF,
		attackPoints,
		defensePoints,
	)
}

func NewWand(name string, attackPoints float64, defensePoints float64) (*Weapon, *errors.InternalError) {
	return new(
		name,
		WAND,
		attackPoints,
		defensePoints,
	)
}

func NewDagger(name string, attackPoints float64) (*Weapon, *errors.InternalError) {
	return new(
		name,
		DAGGER,
		attackPoints,
		0,
	)
}

func (w *Weapon) GetName() string {
	return w.name
}

func (w *Weapon) GetType() WeaponType {
	return w.type_
}

func (w *Weapon) GetBattlePoints() *WeaponAttributes {
	return w.attributes
}

func GetAttackPower(w *Weapon) float64 {
	if w == nil {
		return 0
	}
	return w.GetBattlePoints().GetAttack()
}

func GetDefensePower(w *Weapon) float64 {
	if w == nil {
		return 0
	}
	return w.GetBattlePoints().GetDefense()
}

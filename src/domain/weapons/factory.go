package weapons

import "spanishgab/unind/src/errors"

func NewSword(name string, attackPoints float64, defensePoints float64) (*Weapon, *errors.InternalError) {
	return NewWeapon(
		name,
		Sword,
		attackPoints,
		defensePoints,
	)
}

func NewShield(name string, defensePoints float64) (*Weapon, *errors.InternalError) {
	return NewWeapon(
		name,
		Shield,
		0,
		defensePoints,
	)
}

func NewBowAndArrow(name string, attackPoints float64) (*Weapon, *errors.InternalError) {
	return NewWeapon(
		name,
		BowAndArrow,
		attackPoints,
		0,
	)
}

func NewAxe(name string, attackPoints float64) (*Weapon, *errors.InternalError) {
	return NewWeapon(
		name,
		Axe,
		attackPoints,
		0,
	)
}

func NewStaff(name string, attackPoints float64, defensePoints float64) (*Weapon, *errors.InternalError) {
	return NewWeapon(
		name,
		Staff,
		attackPoints,
		defensePoints,
	)
}

func NewWand(name string, attackPoints float64, defensePoints float64) (*Weapon, *errors.InternalError) {
	return NewWeapon(
		name,
		Wand,
		attackPoints,
		defensePoints,
	)
}

func NewDagger(name string, attackPoints float64) (*Weapon, *errors.InternalError) {
	return NewWeapon(
		name,
		Dagger,
		attackPoints,
		0,
	)
}

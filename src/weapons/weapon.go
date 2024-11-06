package weapons

import (
	"spanishgab/unind/src/attributes"
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
	name         string
	type_        WeaponType
	battlePoints attributes.AttributePoints
}

func New(
	name string,
	type_ WeaponType,
	attackPoints float64,
	defensePoints float64,
) (*Weapon, *errors.InternalError) {
	battlePoints, err := attributes.NewBattlePoints(attackPoints, defensePoints)
	if err != nil {
		return nil, err
	}

	return &Weapon{
		name:         name,
		type_:        type_,
		battlePoints: *battlePoints,
	}, nil
}

func NewSword(name string, attackPoints float64, defensePoints float64) (*Weapon, *errors.InternalError) {
	return New(
		name,
		SWORD,
		attackPoints,
		defensePoints,
	)
}

func NewShield(name string, defensePoints float64) (*Weapon, *errors.InternalError) {
	return New(
		name,
		SHIELD,
		0,
		defensePoints,
	)
}

func NewBowAndArrow(name string, attackPoints float64) (*Weapon, *errors.InternalError) {
	return New(
		name,
		BOW_AND_ARROW,
		attackPoints,
		0,
	)
}

func NewAxe(name string, attackPoints float64) (*Weapon, *errors.InternalError) {
	return New(
		name,
		AXE,
		attackPoints,
		0,
	)
}

func NewStaff(name string, attackPoints float64, defensePoints float64) (*Weapon, *errors.InternalError) {
	return New(
		name,
		STAFF,
		attackPoints,
		defensePoints,
	)
}

func NewWand(name string, attackPoints float64, defensePoints float64) (*Weapon, *errors.InternalError) {
	return New(
		name,
		WAND,
		attackPoints,
		defensePoints,
	)
}

func NewDagger(name string, attackPoints float64) (*Weapon, *errors.InternalError) {
	return New(
		name,
		DAGGER,
		attackPoints,
		0,
	)
}

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
	DAGGERS       WeaponType = "daggers"
)

type Weapon struct {
	name         string
	type_        WeaponType
	battlePoints attributes.AttributePoints
	canAttack    bool
	canDefend    bool
}

// TODO: use flyweight
func NewSword(name string, attackPoints float64, defensePoints float64) (*Weapon, *errors.InternalError) {
	battlePoints, err := attributes.NewBattlePoints(attackPoints, defensePoints)
	if err != nil {
		return nil, err
	}

	return &Weapon{
		name:         name,
		type_:        SWORD,
		battlePoints: *battlePoints,
		canAttack:    true,
		canDefend:    true,
	}, nil
}

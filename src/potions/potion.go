package potions

import (
	"spanishgab/unind/src/errors"
)

type PotionType string

const (
	STRENGTH PotionType = "strength"
	CURE     PotionType = "cure"
)

type Potion struct {
	name          string
	type_         PotionType
	upgradePoints float64
}

// TODO: use flyweight
func NewCurePotion(name string, curePoints float64) (*Potion, *errors.InternalError) {
	return &Potion{
		name:          name,
		type_:         CURE,
		upgradePoints: curePoints,
	}, nil
}

func NewStrengthPotion(name string, strengthPoints float64) (*Potion, *errors.InternalError) {
	return &Potion{
		name:          name,
		type_:         STRENGTH,
		upgradePoints: strengthPoints,
	}, nil
}

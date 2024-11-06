package potions

import (
	"spanishgab/unind/src/errors"
	"spanishgab/unind/src/utils"
)

const (
	INVALID_UPGRADE_VALUE string = "the upgrade value must be greater than zero"
)

type PotionType string

const (
	STRENGTH PotionType = "strength"
	CURE     PotionType = "cure"
)

type Potion struct {
	name          string
	type_         PotionType
	upgradePoints uint
}

func New(name string, type_ PotionType, upgradePoints uint) (*Potion, *errors.InternalError) {
	if !utils.AreAllPositive(upgradePoints) {
		return nil, errors.NewInternalError(INVALID_UPGRADE_VALUE)
	}
	return &Potion{
		name:          name,
		type_:         type_,
		upgradePoints: upgradePoints,
	}, nil
}

func NewCurePotion(name string, curePoints uint) (*Potion, *errors.InternalError) {
	return New(
		name,
		CURE,
		curePoints,
	)
}

func NewStrengthPotion(name string, strengthPoints uint) (*Potion, *errors.InternalError) {
	return New(
		name,
		STRENGTH,
		strengthPoints,
	)
}

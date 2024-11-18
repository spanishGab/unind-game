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

func new(name string, type_ PotionType, upgradePoints uint) (*Potion, *errors.InternalError) {
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
	return new(
		name,
		CURE,
		curePoints,
	)
}

func NewStrengthPotion(name string, strengthPoints uint) (*Potion, *errors.InternalError) {
	return new(
		name,
		STRENGTH,
		strengthPoints,
	)
}

func (p *Potion) GetName() string {
	return p.name
}

func (p *Potion) GetType() PotionType {
	return p.type_
}

func (p *Potion) GetUpgradePoints() uint {
	return p.upgradePoints
}

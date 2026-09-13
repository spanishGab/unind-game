package potions

import (
	"spanishgab/unind/src/errors"
	"spanishgab/unind/src/utils"
)

const (
	InvalidUpgradeValue string = "the upgrade value must be greater than zero"
)

type PotionType string

const (
	Strength PotionType = "strength"
	Cure     PotionType = "cure"
)

type Potion struct {
	name          string
	type_         PotionType
	upgradePoints uint
}

func NewPotion(name string, type_ PotionType, upgradePoints uint) (*Potion, *errors.InternalError) {
	if !utils.AreAllPositive(upgradePoints) {
		return nil, errors.NewInternalError(InvalidUpgradeValue)
	}
	return &Potion{
		name:          name,
		type_:         type_,
		upgradePoints: upgradePoints,
	}, nil
}

func (p *Potion) Name() string {
	return p.name
}

func (p *Potion) Type() PotionType {
	return p.type_
}

func (p *Potion) UpgradePoints() uint {
	return p.upgradePoints
}

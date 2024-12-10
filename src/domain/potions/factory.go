package potions

import "spanishgab/unind/src/errors"

func NewCurePotion(name string, curePoints uint) (*Potion, *errors.InternalError) {
	return NewPotion(
		name,
		Cure,
		curePoints,
	)
}

func NewStrengthPotion(name string, strengthPoints uint) (*Potion, *errors.InternalError) {
	return NewPotion(
		name,
		Strength,
		strengthPoints,
	)
}

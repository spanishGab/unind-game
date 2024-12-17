package creatures

import (
	"spanishgab/unind/src/domain/weapons"
	"spanishgab/unind/src/errors"
)

func NewHuman(
	name string,
	sword *weapons.Weapon,
	shield *weapons.Weapon,
	attributes CreatureAttributes,
) (*Creature, *errors.InternalError) {
	if sword.GetType() != weapons.Sword {
		return nil, MismatchedWeaponTypeError
	}
	if shield.GetType() != weapons.Shield {
		return nil, MismatchedWeaponTypeError
	}
	return NewCreature(name, Human, sword, shield, attributes), nil
}

func NewWizard(
	name string,
	staff *weapons.Weapon,
	sword *weapons.Weapon,
	attributes CreatureAttributes,
) (*Creature, *errors.InternalError) {
	if staff.GetType() != weapons.Staff {
		return nil, MismatchedWeaponTypeError
	}
	if sword.GetType() != weapons.Sword {
		return nil, MismatchedWeaponTypeError
	}
	return NewCreature(name, Wizard, staff, sword, attributes), nil
}

func NewElf(
	name string,
	bowAndArrow *weapons.Weapon,
	attributes CreatureAttributes,
) (*Creature, *errors.InternalError) {
	if bowAndArrow.GetType() != weapons.BowAndArrow {
		return nil, MismatchedWeaponTypeError
	}
	return NewCreature(name, Elf, bowAndArrow, nil, attributes), nil
}

func NewDwarf(
	name string,
	axe *weapons.Weapon,
	attributes CreatureAttributes,
) (*Creature, *errors.InternalError) {
	if axe.GetType() != weapons.Axe {
		return nil, MismatchedWeaponTypeError
	}
	return NewCreature(name, Dwarf, axe, nil, attributes), nil
}

func NewOrc(
	name string,
	sword *weapons.Weapon,
	shield *weapons.Weapon,
	attributes CreatureAttributes,
) (*Creature, *errors.InternalError) {
	if sword.GetType() != weapons.Sword {
		return nil, MismatchedWeaponTypeError
	}
	if shield.GetType() != weapons.Shield {
		return nil, MismatchedWeaponTypeError
	}
	return NewCreature(name, Orc, sword, shield, attributes), nil
}

func NewSorcerer(
	name string,
	wand *weapons.Weapon,
	attributes CreatureAttributes,
) (*Creature, *errors.InternalError) {
	if wand.GetType() != weapons.Wand {
		return nil, MismatchedWeaponTypeError
	}
	return NewCreature(name, Orc, wand, nil, attributes), nil
}

func NewShadowLord(
	name string,
	staff *weapons.Weapon,
	sword *weapons.Weapon,
	attributes CreatureAttributes,
) (*Creature, *errors.InternalError) {
	if staff.GetType() != weapons.Staff {
		return nil, MismatchedWeaponTypeError
	}
	if sword.GetType() != weapons.Sword {
		return nil, MismatchedWeaponTypeError
	}
	return NewCreature(name, Orc, staff, sword, attributes), nil
}

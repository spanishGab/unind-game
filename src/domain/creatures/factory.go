package creatures

import "spanishgab/unind/src/domain/weapons"

func NewHuman(name string) *Creature {
	sword, _ := weapons.NewSword("Iron sword", 7, 2)
	shield, _ := weapons.NewShield("Wooden shield", 7)
	attributes, _ := NewCreatureAttributes(75, 85, 80)
	return NewCreature(name, Human, sword, shield, *attributes)
}

func NewWizard(name string) *Creature {
	sword, _ := weapons.NewStaff("Wooden staff", 10, 10)
	attributes, _ := NewCreatureAttributes(90, 95, 85)
	return NewCreature(name, Human, sword, nil, *attributes)
}

func NewElf(name string) *Creature {
	sword, _ := weapons.NewStaff("Wooden staff", 10, 10)
	attributes, _ := NewCreatureAttributes(90, 95, 85)
	return NewCreature(name, Human, sword, nil, *attributes)
}

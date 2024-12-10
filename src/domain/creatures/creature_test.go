package creatures

import (
	"fmt"
	"spanishgab/unind/src/domain/potions"
	"spanishgab/unind/src/domain/weapons"
	"spanishgab/unind/src/errors"
	"testing"
)

func TestAttack(t *testing.T) {
	falameShot, _ := weapons.NewDagger("Flame shot", 7)
	icyShot, _ := weapons.NewDagger("Icy shot", 6)
	attributePoints, _ := NewCreatureAttributes(90, 70, 85)
	elf := NewCreature(
		"Legolas",
		Elf,
		falameShot,
		icyShot,
		*attributePoints,
	)

	t.Run("Should return the attack points successfully", func(t *testing.T) {
		var expected float64 = 69
		got := elf.Attack()

		if got != expected {
			t.Errorf("expected: %f, got: %f", expected, got)
		}
	})
}

func TestDefend(t *testing.T) {
	woodenShield, _ := weapons.NewShield("Wooden shield", 10)
	ironSword, _ := weapons.NewSword("Iron sword", 8, 2)
	attributePoints, _ := NewCreatureAttributes(90, 70, 85)
	elf := NewCreature(
		"Legolas",
		Elf,
		woodenShield,
		ironSword,
		*attributePoints,
	)

	t.Run("Should reduce the health points when the attack received is grater thant the defense points", func(t *testing.T) {
		var expected float64 = 88
		var attack float64 = 70
		elf.Defend(attack)

		if elf.attributes.GetHealth() != expected {
			t.Errorf("expected: %f, got: %f", expected, elf.attributes.GetHealth())
		}
	})

	t.Run("Should reduce the creature's health to zero and return a DomainError", func(t *testing.T) {
		var expectedHealth float64 = 0
		var expectedError *errors.DomainError = CreatureDiedError
		var attack float64 = 200
		err := elf.Defend(attack)

		if elf.attributes.GetHealth() != expectedHealth {
			t.Errorf("expected: %f, got: %f", expectedHealth, elf.attributes.GetHealth())
		}
		if !err.Is(expectedError) {
			t.Errorf("expected: %v, got: %v", expectedError, err)
		}
	})
}

func TestHeal(t *testing.T) {
	healingPotion50Points, _ := potions.NewCurePotion("Healing potion", 50)
	healingPotion5Points, _ := potions.NewCurePotion("Healing potion", 5)

	ironAxe, _ := weapons.NewAxe("Iron axe", 9)
	attributePoints, _ := NewCreatureAttributes(90, 70, 85)
	dwarf := NewCreature(
		"Gimli",
		Dwarf,
		nil,
		ironAxe,
		*attributePoints,
	)

	params := []struct {
		should   string
		potion   *potions.Potion
		expected float64
	}{
		{
			should:   "Should upgrade the creature health points unitl it hits the maximum health points allowed",
			potion:   healingPotion50Points,
			expected: MaxHealthPoints,
		},
		{
			should:   "Should heal the creature based on the potion's curing points",
			potion:   healingPotion5Points,
			expected: 95,
		},
	}
	for _, param := range params {
		t.Run(fmt.Sprintf("Should %s", param.should), func(t *testing.T) {
			t.Cleanup(func() { dwarf.attributes.SetHealth(90) })

			dwarf.Heal(param.potion)

			if dwarf.attributes.GetHealth() != param.expected {
				t.Errorf("expected: %f, got: %f", param.expected, dwarf.attributes.GetHealth())
			}
		})
	}
}

func TestGetBattlePoints(t *testing.T) {
	woodenShield, _ := weapons.NewShield("Wooden shield", 10)
	ironSword, _ := weapons.NewSword("Iron sword", 8, 2)
	attributePoints, _ := NewCreatureAttributes(90, 70, 85)
	elf := NewCreature(
		"Legolas",
		Elf,
		woodenShield,
		ironSword,
		*attributePoints,
	)

	t.Run("Should return the battle points successfully, no matter if attack or defense points", func(t *testing.T) {
		var expected float64 = 68
		got := elf.getBattlePoints(12)

		if got != expected {
			t.Errorf("expected: %f, got: %f", expected, got)
		}
	})
}

func TestStrengthen(t *testing.T) {
	woodenStaff, _ := weapons.NewStaff("Wooden staff", 15, 15)
	attributePoints, _ := NewCreatureAttributes(90, 100, 80)
	wizard := NewCreature("Gandalf", Wizard, woodenStaff, nil, *attributePoints)
	strengthenPotion19Points, _ := potions.NewStrengthPotion("Young warrior potion", 19)
	strengthenPotion21Points, _ := potions.NewStrengthPotion("Young warrior potion", 21)

	params := []struct {
		should   string
		potion   *potions.Potion
		expected float64
	}{
		{
			should:   "Should upgrade the creature strenght points unitl it hits the maximum strength points allowed",
			potion:   strengthenPotion21Points,
			expected: MaxStrengthPoints,
		},
		{
			should:   "Should strengthen the creature according to the potion's strength points",
			potion:   strengthenPotion19Points,
			expected: 99,
		},
	}
	for _, param := range params {
		t.Run(fmt.Sprintf("Should %s", param.should), func(t *testing.T) {
			t.Cleanup(func() { wizard.attributes.SetStrength(80) })

			wizard.Strengthen(param.potion)

			if wizard.attributes.GetStrength() != param.expected {
				t.Errorf("expected: %f, got: %f", param.expected, wizard.attributes.GetStrength())
			}
		})
	}
}

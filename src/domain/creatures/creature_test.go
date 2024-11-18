package creatures

import (
	"spanishgab/unind/src/domain/potions"
	"spanishgab/unind/src/domain/weapons"
	"testing"
)

func TestAttack(t *testing.T) {
	falameShot, _ := weapons.NewDagger("Flame shot", 7)
	icyShot, _ := weapons.NewDagger("Icy shot", 6)
	attributePoints, _ := NewCreatureAttributes(90, 70, 85)
	elf := New(
		"Legolas",
		ELF,
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
	elf := New(
		"Legolas",
		ELF,
		woodenShield,
		ironSword,
		*attributePoints,
	)

	t.Run("Should return the defense points successfully", func(t *testing.T) {
		var expected float64 = 68
		got := elf.Defend()

		if got != expected {
			t.Errorf("expected: %f, got: %f", expected, got)
		}
	})
}

func TestHeal(t *testing.T) {
	ironAxe, _ := weapons.NewAxe("Iron axe", 9)
	attributePoints, _ := NewCreatureAttributes(90, 70, 85)
	dwarf := New(
		"Gimli",
		DWARF,
		nil,
		ironAxe,
		*attributePoints,
	)

	t.Run("Should upgrade the creature health points unitl it hits the maximum health points allowed", func(t *testing.T) {
		t.Cleanup(func() { dwarf.attributes.SetHealth(90) })

		healingPotion, _ := potions.NewCurePotion("Healing potion", 50)
		var expected float64 = MAX_HEALTH_POINTS
		dwarf.Heal(healingPotion)

		if dwarf.attributes.health != expected {
			t.Errorf("expected: %f, got: %f", expected, dwarf.attributes.health)
		}
	})
	t.Run("Should heal the creature based on the potion curing points", func(t *testing.T) {
		t.Cleanup(func() { dwarf.attributes.SetHealth(90) })

		healingPotion, _ := potions.NewCurePotion("Healing potion", 5)
		var expected float64 = 95
		dwarf.Heal(healingPotion)

		if dwarf.attributes.health != expected {
			t.Errorf("expected: %f, got: %f", expected, dwarf.attributes.health)
		}
	})
}

func TestGetBattlePoints(t *testing.T) {
	woodenShield, _ := weapons.NewShield("Wooden shield", 10)
	ironSword, _ := weapons.NewSword("Iron sword", 8, 2)
	attributePoints, _ := NewCreatureAttributes(90, 70, 85)
	elf := New(
		"Legolas",
		ELF,
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

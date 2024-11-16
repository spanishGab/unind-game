package creatures

import (
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

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
		var expected float64 = 68.6
		got := elf.Attack()

		if got != expected {
			t.Errorf("expected: %.2f, got: %.2f", expected, got)
		}
	})

	t.Run("Should return the attack points successfully", func(t *testing.T) {
		var expected float64 = 68.6
		got := elf.Attack()

		if got != expected {
			t.Errorf("expected: %.2f, got: %.2f", expected, got)
		}
	})
}

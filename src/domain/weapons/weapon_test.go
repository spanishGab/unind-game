package weapons

import "testing"

func TestGetAttackPower(t *testing.T) {
	t.Run("Should return the attack points of a weapon when it has attack power", func(t *testing.T) {
		ironSword, _ := NewSword("Iron sword", 8, 2)

		var expected float64 = 8
		got := GetAttackPower(ironSword)

		if got != expected {
			t.Errorf("expected: %f, got: %f", expected, got)
		}
	})

	t.Run("Should return 0 when the weapon has not attack power", func(t *testing.T) {
		woodenShield, _ := NewShield("Wooden shield", 10)

		var expected float64 = 0
		got := GetAttackPower(woodenShield)

		if got != expected {
			t.Errorf("expected: %f, got: %f", expected, got)
		}
	})
}

func TestGetDefensePower(t *testing.T) {
	t.Run("Should return the defense points of a weapon when it has defense power", func(t *testing.T) {
		woodenShield, _ := NewShield("Wooden shield", 10)

		var expected float64 = 10
		got := GetDefensePower(woodenShield)

		if got != expected {
			t.Errorf("expected: %f, got: %f", expected, got)
		}
	})

	t.Run("Should return 0 when the weapon has not attack power", func(t *testing.T) {
		falameShot, _ := NewDagger("Flame shot", 7)

		var expected float64 = 0
		got := GetDefensePower(falameShot)

		if got != expected {
			t.Errorf("expected: %f, got: %f", expected, got)
		}
	})
}

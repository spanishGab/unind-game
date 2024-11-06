package utils

import "testing"

func TestArePositiveAttributeValues(t *testing.T) {
	t.Run("Should return an error when any of the given values are less than zero", func(t *testing.T) {
		expected := false
		got := AreAllPositive(30, 25, -1)

		if got != expected {
			t.Errorf("expected: %t, got: %t", expected, got)
		}
	})

	t.Run("Should return nil all of the given values are greater than or equal to zero", func(t *testing.T) {
		expected := true
		got := AreAllPositive(30, 25, 0)

		if got != expected {
			t.Errorf("expected: %t, got: %t", expected, got)
		}
	})
}

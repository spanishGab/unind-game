package attributes

import (
	"fmt"
	"testing"
)

func createInternalErrorMessage(errorMessage string, t *testing.T) string {
	t.Helper()
	return fmt.Sprintf("internal-error: %s", errorMessage)
}

func TestGet(t *testing.T) {
	attributePoints, _ := NewBattlePoints(45, 0)
	t.Run("Should return a value for an existing key", func(t *testing.T) {
		var expected float64 = 45.0
		got, _ := attributePoints.Get(ATTACK)

		if got != expected {
			t.Errorf("expected: %g, got: %g", expected, got)
		}
	})

	t.Run("Should return an error when the given key does not exist", func(t *testing.T) {
		expected := createInternalErrorMessage(fmt.Sprintf(ATTRIBUTE_NOT_FOUND, HEALTH), t)
		_, got := attributePoints.Get(HEALTH)

		if got.Error() != expected {
			t.Errorf("expected: %q, got: %q", expected, got.Error())
		}

	})
}

func TestHas(t *testing.T) {
	attributePoints, _ := NewBattlePoints(50, 20)

	params := []struct {
		should    string
		attribute string
		expected  bool
	}{
		{
			should:    "return true when the given attribute exists",
			attribute: string(DEFENSE),
			expected:  true,
		},
		{
			should:    "return false when the given attribute does not exist",
			attribute: "potato",
			expected:  false,
		},
	}
	for _, param := range params {
		t.Run(fmt.Sprintf("Should %s", param.should), func(t *testing.T) {
			got := attributePoints.Has(Attribute(param.attribute))

			if got != param.expected {
				t.Errorf("expected %t, got %t", param.expected, got)
			}
		})
	}
}

func TestIncrease(t *testing.T) {
	attributePoints, _ := NewBattlePoints(45, 62)
	t.Run("Should increase an attribute's value", func(t *testing.T) {
		var expected float64 = 65.0
		attributePoints.Increase(DEFENSE, 3)
		got, _ := attributePoints.Get(DEFENSE)

		if got != expected {
			t.Errorf("expected: %g, got %g", expected, got)
		}
	})

	t.Run("Should return an error when trying to increase an inexistent attribute", func(t *testing.T) {
		expected := createInternalErrorMessage(fmt.Sprintf(ATTRIBUTE_NOT_FOUND, HEALTH), t)
		got := attributePoints.Increase(HEALTH, 3)

		if got.Error() != expected {
			t.Errorf("expected: %q, got %q", expected, got.Error())
		}
	})

	t.Run("Should return an error when the increasing points are less than zero", func(t *testing.T) {
		expected := createInternalErrorMessage(INVALID_INCREASING_VALUE, t)
		got := attributePoints.Increase(ATTACK, -3)

		if got.Error() != expected {
			t.Errorf("expected: %q, got %q", expected, got.Error())
		}
	})
}

func TestDecrease(t *testing.T) {
	attributePoints, _ := NewBattlePoints(45, 62)
	t.Run("Should decrease an attribute's value", func(t *testing.T) {
		var expected float64 = 55.0
		attributePoints.Decrease(DEFENSE, -7)
		got, _ := attributePoints.Get(DEFENSE)

		if got != expected {
			t.Errorf("expected: %g, got %g", expected, got)
		}
	})

	t.Run("Should return an error when trying to decrease an inexistent attribute", func(t *testing.T) {
		expected := createInternalErrorMessage(fmt.Sprintf(ATTRIBUTE_NOT_FOUND, HEALTH), t)
		got := attributePoints.Decrease(HEALTH, 3)

		if got.Error() != expected {
			t.Errorf("expected: %q, got %q", expected, got.Error())
		}
	})

	t.Run("Should return an error when the decreasing points are greater than zero", func(t *testing.T) {
		expected := createInternalErrorMessage(INVALID_DECREASING_VALUE, t)
		got := attributePoints.Decrease(ATTACK, 3)

		if got.Error() != expected {
			t.Errorf("expected: %q, got %q", expected, got.Error())
		}
	})
}

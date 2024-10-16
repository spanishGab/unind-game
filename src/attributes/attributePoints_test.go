package attributes

import (
	"fmt"
	"spanishgab/unind/src/errors"
	"testing"
)

func TestGet(t *testing.T) {
	attributePoints := AttributePoints{
		points: map[Attribute]float64{
			ATTACK: 45,
		},
	}
	t.Run("Should return a value for an existing key", func(t *testing.T) {
		var expected float64 = 45.0
		got, _ := attributePoints.Get(ATTACK)

		if got != expected {
			t.Errorf("expected: %g, got: %g", expected, got)
		}
	})

	t.Run("Should return an error when the given key does not exist", func(t *testing.T) {
		var expected *errors.InternalError = errors.NewInternalError(
			fmt.Sprintf(ATTRIBUTE_NOT_FOUND, DEFENSE),
		)
		_, got := attributePoints.Get(DEFENSE)

		if got.Error() != expected.Error() {
			t.Errorf("expected: %q, got: %q", expected.Error(), got.Error())
		}

	})
}

func TestHas(t *testing.T) {
	attributePoints := AttributePoints{
		points: map[Attribute]float64{
			ATTACK: 45,
			DEFENSE: 62,
			INTELLIGENCE: 50,
		},
	}
	
	params := []struct{
		should string
		attribute string
		expected bool
	}{
		{
			should: "return true when the given attribute exists",
			attribute: string(INTELLIGENCE),
			expected: true,
		},
		{
			should: "return false when the given attribute does not exist",
			attribute: "potato",
			expected: false,
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
	attributePoints := AttributePoints{
		points: map[Attribute]float64{
			ATTACK: 45,
			DEFENSE: 62,
		},
	}
	t.Run("Should increase an attribute's value", func(t *testing.T) {
		var expected float64 = 65.0
		attributePoints.Increase(DEFENSE, 3)
		got, _ := attributePoints.Get(DEFENSE)

		if got != expected {
			t.Errorf("expected: %g, got %g", expected, got)
		}
	})

	t.Run("Should return an error when trying to increase an inexistent attribute", func(t *testing.T) {
		var expected *errors.InternalError = errors.NewInternalError(
			fmt.Sprintf(ATTRIBUTE_NOT_FOUND, HEALTH),
		)
		got := attributePoints.Increase(HEALTH, 3)

		if got.Error() != expected.Error() {
			t.Errorf("expected: %q, got %q", expected.Error(), got.Error())
		}
	})

	t.Run("Should return an error when the increasing points are less than zero", func(t *testing.T) {
		var expected *errors.InternalError = errors.NewInternalError(
			INVALID_INCREASING_VALUE,
		)
		got := attributePoints.Increase(ATTACK, -3)

		if got.Error() != expected.Error() {
			t.Errorf("expected: %q, got %q", expected.Error(), got.Error())
		}
	})
}

func TestDecrease(t *testing.T) {
	attributePoints := AttributePoints{
		points: map[Attribute]float64{
			ATTACK: 45,
			DEFENSE: 62,
		},
	}
	t.Run("Should decrease an attribute's value", func(t *testing.T) {
		var expected float64 = 55.0
		attributePoints.Decrease(DEFENSE, -7)
		got, _ := attributePoints.Get(DEFENSE)

		if got != expected {
			t.Errorf("expected: %g, got %g", expected, got)
		}
	})

	t.Run("Should return an error when trying to decrease an inexistent attribute", func(t *testing.T) {
		var expected *errors.InternalError = errors.NewInternalError(
			fmt.Sprintf(ATTRIBUTE_NOT_FOUND, HEALTH),
		)
		got := attributePoints.Decrease(HEALTH, 3)

		if got.Error() != expected.Error() {
			t.Errorf("expected: %q, got %q", expected.Error(), got.Error())
		}
	})

	t.Run("Should return an error when the decreasing points are greater than zero", func(t *testing.T) {
		var expected *errors.InternalError = errors.NewInternalError(
			INVALID_DECREASING_VALUE,
		)
		got := attributePoints.Decrease(ATTACK, 3)

		if got.Error() != expected.Error() {
			t.Errorf("expected: %q, got %q", expected.Error(), got.Error())
		}
	})
}
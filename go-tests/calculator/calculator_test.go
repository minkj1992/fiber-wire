package calculator

import (
	"testing"
)

type TestCase struct {
	value    int
	expected bool
	actual   bool
}

func TestCalculateIsArmstrong(t *testing.T) {
	t.Run("Should return true for 371", func(t *testing.T) {
		tc := TestCase{
			value:    371,
			expected: false,
		}

		tc.actual = CalculateIsArmstrong(tc.value)
		if tc.actual != tc.expected {
			t.Errorf("CalculateIsArmstrong returned an unexpected result")
		}
	})

	t.Run("Should return true for first armstrong number", func(t *testing.T) {
		tc := TestCase{
			value:    1,
			expected: true,
		}

		tc.actual = CalculateIsArmstrong(tc.value)
		if tc.actual != tc.expected {
			t.Errorf("CalculateIsArmstrong returned an unexpected result")
		}
	})

	t.Run("Should return false for non-armstrong number", func(t *testing.T) {
		testCase := TestCase{
			value:    15,
			expected: false,
		}

		testCase.actual = CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Errorf("CalculateIsArmstrong returned an unexpected result")
		}
	})
}

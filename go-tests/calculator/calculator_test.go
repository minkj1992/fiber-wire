package calculator

import (
	"testing"
)

type TestCase struct {
	name     string
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

func TestCalculateIsArmstrongWithTable(t *testing.T) {
	t.Run("test for all 3 digit armstrong numbers", func(t *testing.T) {
		tests := []TestCase{
			{name: "Testing value for: 153", value: 153, expected: true},
			{name: "Testing value for: 370", value: 370, expected: true},
			{name: "Testing value for: 371", value: 371, expected: true},
			{name: "Testing value for: 407", value: 407, expected: true},
		}

		for _, test := range tests {
			t.Run("test", func(t *testing.T) {
				actual := CalculateIsArmstrong(test.value)
				if test.expected != actual {
					t.Fail()
				}
			})
		}
	})

}

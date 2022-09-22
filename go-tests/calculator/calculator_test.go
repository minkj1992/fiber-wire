package calculator_test

import (
	"testing"

	"github.com/minkj1992/go-playground/go-tests/calculator"
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
			expected: true,
		}

		tc.actual = calculator.CalculateIsArmstrong(tc.value)
		if tc.actual != tc.expected {
			t.Errorf("CalculateIsArmstrong returned an unexpected result")
		}
	})

	t.Run("Should return true for first armstrong number", func(t *testing.T) {
		tc := TestCase{
			value:    1,
			expected: true,
		}

		tc.actual = calculator.CalculateIsArmstrong(tc.value)
		if tc.actual != tc.expected {
			t.Errorf("CalculateIsArmstrong returned an unexpected result")
		}
	})

	t.Run("Should return false for non-armstrong number", func(t *testing.T) {
		testCase := TestCase{
			value:    15,
			expected: false,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
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
				actual := calculator.CalculateIsArmstrong(test.value)
				if test.expected != actual {
					t.Fail()
				}
			})
		}
	})
}

func benchmarkCalculateIsArmstrong(input int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculator.CalculateIsArmstrong(input)
	}
}

func BenchmarkCalculateIsArmstrong370(b *testing.B) { benchmarkCalculateIsArmstrong(370, b) }
func BenchmarkCalculateIsArmstrong371(b *testing.B) { benchmarkCalculateIsArmstrong(371, b) }
func BenchmarkCalculateIsArmstrong0(b *testing.B)   { benchmarkCalculateIsArmstrong(0, b) }

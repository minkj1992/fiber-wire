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
	tc := TestCase{
		value:    371,
		expected: true,
	}

	tc.actual = CalculateIsArmstrong(tc.value)
	if tc.actual != tc.expected {
		t.Fail()
	}
}

func TestNegativeCalculateIsArmstrong(t *testing.T) {
	testCase := TestCase{
		value:    372,
		expected: false,
	}

	testCase.actual = CalculateIsArmstrong(testCase.value)
	if testCase.actual != testCase.expected {
		t.Fail()
	}
}

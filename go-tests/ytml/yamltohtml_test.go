package ytml_test

import (
	"testing"

	"github.com/minkj1992/go-playground/go-tests/ytml"
)

type TestCase struct {
	desc     string
	path     string
	expected string
}

func TestYAMLToHTML(t *testing.T) {
	testCases := []TestCase{
		{
			desc:     "Tests title is set properly",
			path:     "testdata/ytml_01.yml",
			expected: "<html><head><title>My Awesome Page</title></head><body>This is my awesome body</body></html>",
		},
		{
			desc:     "Tests body is set properly",
			path:     "testdata/ytml_02.yml",
			expected: "<html><head><title>Second Page</title></head><body>This is a second page</body></html>",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			result, err := ytml.YamlToHTML(test.path)
			if err != nil {
				t.Errorf("Invalid path is given.")
			}

			t.Log(result)

			if result != test.expected {
				t.Errorf("The result is not matched with expected value.")
			}
		})
	}

}

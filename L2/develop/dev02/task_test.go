package main

//to launch this test:
//cd dev02
//go test
//You can add the -v flag to get verbose output that lists all of the tests and their results, so "go test -v"

import (
	"errors"
	"testing"
)

func TestUnpackString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		err      error
	}{ // success cases
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"", "", nil},
		{"qwe\\4\\5", "qwe45", nil},
		{"qwe\\45", "qwe44444", nil},
		{"qwe\\\\5", "qwe\\\\\\\\\\", nil},

		// error cases
		{"45", "", errors.New("некорректная строка")},
		{"qwe\\\\5\\", "", errors.New("некорректная строка")}}

	for i, test := range testCases {
		t.Run(test.input, func(t *testing.T) {
			result, err := unpackString(test.input)
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
			if err != nil && test.err != nil && err.Error() != test.err.Error() {
				t.Errorf("Test %d: expected %v, but got %v", i, test.err, err)
			}
		})
	}
}

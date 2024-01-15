package main

import (
	"errors"
	"testing"
)

func Test(t *testing.T) {
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

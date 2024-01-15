package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func unpackString(str string) (string, error) {
	// use strings.Builder as there are a lot of concatenations with string
	var res strings.Builder
	s := []rune(str)
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(s[0]) {
			return "", errors.New("некорректная строка")
		}
		if s[len(s)-1] == '\\' {
			return "", errors.New("некорректная строка")
		}
		if s[i] == '\\' {
			res.WriteString(string(s[i+1]))
			i++
		} else {
			if unicode.IsDigit(s[i]) {
				count, _ := strconv.Atoi(string(s[i]))
				res.WriteString(strings.Repeat(string(s[i-1]), count-1))

			} else if i < len(s)-1 && !unicode.IsDigit(s[i+1]) {
				res.WriteString(string(s[i]))
			} else {
				res.WriteString(string(s[i]))
			}
		}
	}
	return res.String(), nil
}

func main() {
	s := "a4bc2d5e"
	// unpack string
	result, err := unpackString(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

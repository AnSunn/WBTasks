package task26

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	m := make(map[rune]bool)
	s = strings.ToLower(s)
	for _, val := range s {
		if m[val] {
			return false
		}
		m[val] = true
	}
	return true
}

func Launch() {
	str := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, v := range str {
		fmt.Println(v, isUnique(v))
	}
}

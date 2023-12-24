package task19

import (
	"fmt"
)

func Launch() {
	word := "главрыба"
	fmt.Println("The reversed string is: ", reverseWord(word))
}

func reverseWord(s string) string {
	chars := []rune(s) //type rune = int32
	for i := 0; i < len(chars)/2; i++ {
		chars[i], chars[len(chars)-i-1] = chars[len(chars)-i-1], chars[i]
	}
	return string(chars)
}

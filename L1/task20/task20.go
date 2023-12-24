package task20

import (
	"fmt"
	"strings"
)

func Launch() {
	str := "snow dog sun"
	words := strings.Split(str, " ")
	length := len(words)
	for i := 0; i < length/2; i++ {
		words[i], words[length-i-1] = words[length-i-1], words[i]
	}
	fmt.Println(strings.Join(words, " "))
}

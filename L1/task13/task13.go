package task13

import "fmt"

func Launch() {
	a, b := 10, 20
	a, b = b, a
	fmt.Println(a, b)
}

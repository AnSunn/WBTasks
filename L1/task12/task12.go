package task12

import "fmt"

func Launch() {

	val := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]struct{})
	for _, j := range val {
		m[j] = struct{}{}
	}
	fmt.Println("The result is: ", m)
}

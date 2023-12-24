package task15

import (
	"fmt"
	"math/rand"
	"runtime"
)

//A potential memory leak is possible because a large string is created in createHugeString(1 << 10),
//and then only its first 100 bytes are copied into justString. However, the memory allocated for
//the entire string is not freed. This results in the remaining part of the string becoming inaccessible
//for use, yet it is not returned to the system for reuse.

var justString string

func Launch() {
	someFunc()
	fmt.Println(justString)

	//Call GC for explicit memory release
	runtime.GC()
}
func createHugeString(n int) string {
	const letters = "longstring"
	res := make([]byte, n)
	for i := range res {
		res[i] = letters[rand.Intn(len(letters))]
	}
	return string(res)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
	//Free the memory allocated for the 'v' string
	v = ""
}

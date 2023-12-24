package task23

import "fmt"

func Launch() {
	sl := []int{1, 2, 4, 6, 8}
	removeElement(&sl, 3)
	fmt.Println(sl)
}

func removeElement(sl *[]int, elem int) {
	*sl = append((*sl)[:elem], (*sl)[elem+1:]...)
}

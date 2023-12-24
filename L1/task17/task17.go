package task17

import (
	"fmt"
	"sort"
)

func Launch() {
	arr := []float64{3.2, 1.5, 4.8, 2.1, 5.8}
	// sort slice
	sort.Float64s(arr)
	// lookingValue
	lookingValue := 5.80
	// Use built-in method. For int slice use sort.SearchInts()
	index := sort.SearchFloat64s(arr, lookingValue)
	// Check is the value found
	if index < len(arr) && arr[index] == lookingValue {
		fmt.Printf("Value %.2f found at index %d\n", lookingValue, index)
	} else {
		fmt.Printf("Value %.2f not found in the array\n", lookingValue)
	}
}

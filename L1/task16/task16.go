package task16

import (
	"fmt"
	"sort"
)

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivotIndex := partition(arr)

	quickSort(arr[:pivotIndex])
	quickSort(arr[pivotIndex+1:])
}

func quickSort2(arr []int) {
	// Check the length of array
	if len(arr) <= 1 {
		return
	}

	// Use go-methods
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

func partition(arr []int) int {
	pivotIndex := len(arr) / 2
	pivotValue := arr[pivotIndex]

	left := 0
	right := len(arr) - 1

	for left <= right {
		for arr[left] < pivotValue {
			left++
		}

		for arr[right] > pivotValue {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	return left - 1
}
func Launch() {
	arr := []int{5, 2, 9, 10, 5, 6}
	quickSort(arr)
	fmt.Println("Option one: Sorted array:", arr)

	arr2 := []int{5, 23, 49, 10, 5, 63}
	quickSort2(arr2)
	fmt.Println("Option two: Sorted array:", arr2)
}

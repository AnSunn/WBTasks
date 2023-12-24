package task11

import "fmt"

func Launch() {
	set1 := make(map[int64]struct{})
	set2 := make(map[int64]struct{})
	keys1 := []int64{2, 4, 6, 8}
	keys2 := []int64{4, 6, 7, 8, 10}
	fillInValues(keys1, set1)
	fillInValues(keys2, set2)
	res := intersection(set1, set2)
	for _, j := range res {
		fmt.Println("The result of sets intersection is ", j)
	}

}

func fillInValues(keys []int64, m map[int64]struct{}) {
	for _, key := range keys {
		m[key] = struct{}{}
	}
}

func intersection(set1 map[int64]struct{}, set2 map[int64]struct{}) []int64 {
	duplicates := make(map[int64]bool)
	res := []int64{}
	for i, _ := range set1 {
		duplicates[i] = true
	}

	for i, _ := range set2 {
		if duplicates[i] {
			res = append(res, i)
		}
	}
	return res
}

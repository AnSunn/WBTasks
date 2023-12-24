package task9

import (
	"fmt"
	"sync"
)

func Launch() {
	array := [5]int64{2, 4, 6, 8, 10}
	chanValues := make(chan int64)
	chanDoubleValues := make(chan int64)
	m := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go fillInValues(array, chanValues, m, wg)
	go fillInDoubleValues(wg, chanValues, chanDoubleValues)
	for i := range chanDoubleValues {
		fmt.Println(i)
	}
	wg.Wait()
}

func fillInValues(arr [5]int64, cw chan int64, m *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, j := range arr {
		cw <- j
	}
	close(cw)
}
func fillInDoubleValues(wg *sync.WaitGroup, chanValues chan int64, cr chan int64) {
	defer wg.Done()
	for i := range chanValues {
		cr <- i * 2
	}
	close(cr)
}

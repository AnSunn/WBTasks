package task18

import (
	"fmt"
	"sync"
)

type increment struct {
	count int64
	m     sync.Mutex
}

func Launch() {
	var val = increment{count: 0}
	var wg sync.WaitGroup
	n := 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			val.m.Lock()
			val.count++
			val.m.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(val.count)
}

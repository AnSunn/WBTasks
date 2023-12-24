package task7

import (
	"fmt"
	"sync"
)

func option1() {
	m := make(map[int]int)
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i, wg, mu)
	}
	wg.Wait()
	fmt.Println(m)
}

// Use of channel
func option2() {
	m := make(map[int]int)
	wg := new(sync.WaitGroup)
	type DataSafe struct {
		Key int
		Val int
	}
	c := make(chan DataSafe)
	go func() {
		for v := range c {
			m[v.Key] = v.Val
		}
		wg.Done()
	}()
	wg.Add(1)
	c <- DataSafe{Key: 0, Val: 1}
	c <- DataSafe{Key: 1, Val: 2}
	c <- DataSafe{Key: 2, Val: 3}
	c <- DataSafe{Key: 3, Val: 4}
	close(c)
	wg.Wait()
	fmt.Println(m)
}

func Launch() {
	fmt.Println("option 1 using mutex")
	option1()

	fmt.Println("option 2 using channel")
	option2()
}

package task2

import (
	"fmt"
	"math"
)

func Launch() {
	array := [5]int{2, 4, 6, 8, 10}
	c := make(chan struct{})
	//wg := new(sync.WaitGroup)
	for _, j := range array {
		//	wg.Add(1)
		//go squared(j, wg, c)
		go squares(j, c)
		c <- struct{}{}
	}
	//wg.Wait()
	defer close(c)

}

func squares(j int, c chan struct{}) {
	//(j int, wg *sync.WaitGroup, c chan struct{}) {
	//defer wg.Done()
	fmt.Println(int(math.Pow(float64(j), 2)))
	<-c
}

func squares2(arr [5]int, c chan int) {
	for _, j := range arr {
		c <- j * j
	}
	close(c)
}

func Launch2() {
	array := [5]int{2, 4, 6, 8, 10}
	c := make(chan int)
	//wg := new(sync.WaitGroup)
	go squares2(array, c)
	for val := range c {
		fmt.Println(val)
	}

}

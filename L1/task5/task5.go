package task5

import (
	"fmt"
	"time"
)

func Launch() {
	c := make(chan int)
	n := 5
	timeout := time.After(time.Duration(n) * time.Second)
	for i := 0; i < 10; i++ {
		go fillInChan(c, i)
		fmt.Println(<-c, " value")
	}
	for {
		select {
		case <-timeout:
			fmt.Println("the time is over")
			close(c)
			return
		}
	}
}
func fillInChan(c chan int, i int) {
	c <- i
}

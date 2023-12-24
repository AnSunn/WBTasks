package task6

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func option1() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println("goroutine is running")
		}(wg)
	}
	wg.Wait()
	fmt.Println("main func is over")

}

func Launch() {
	fmt.Println("option 1 using wait groups:")
	option1()
	fmt.Println("option 2 using channels:")
	option2()
	fmt.Println("option 3 using closure of channel:")
	option3()
	fmt.Println("option 4 using context:")
	option4()
}

func option2() {
	c := make(chan struct{})
	for i := 0; i < 5; i++ {
		go func(c chan struct{}) {
			fmt.Println("goroutine is running")
			c <- struct{}{}
		}(c)
		<-c
	}
}

func option3() {
	stopCh := make(chan struct{})
	go func() {
		for {
			select {
			case <-stopCh:
				fmt.Println("the time is over")
				return
			default:
				time.Sleep(1 * time.Second)
				fmt.Println("goroutine is running")
			}

		}
	}()

	time.Sleep(2 * time.Second)
	close(stopCh)
	//Wait for several seconds to print the over message
	time.Sleep(3 * time.Second)
}

func option4() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("time is over")
				return
			default:
				fmt.Println("goroutine is running")
			}
		}
	}()

	//Wait for goroutine to print the over message
	time.Sleep(1 * time.Millisecond)
}

package task3

import "fmt"

func Launch() {
	arr := [5]int{2, 4, 6, 8, 10}
	sum := 0
	c := make(chan int)
	go squares(arr, c)
	for val := range c {
		sum += val
	}
	fmt.Println(sum)
}

func squares(arr [5]int, c chan int) {
	for _, j := range arr {
		c <- j * j
	}
	close(c)

}

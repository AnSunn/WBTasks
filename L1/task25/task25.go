package task25

import (
	"fmt"
	"time"
)

func sleep() {
	<-time.After(time.Second * 7)
}

func Launch() {
	fmt.Println("Go to sleep...")
	sleep()
	fmt.Println("Woke up")
}

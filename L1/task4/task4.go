package task4

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	for {
		select {
		case data, ok := <-ch:
			if !ok {
				fmt.Printf("Worker %d shutting down\n", id)
				return
			}
			fmt.Printf("Worker %d received: %d\n", id, data)
		}
	}
}

func Launch() {
	const numWorkers = 3 //Number of workers

	dataChannel := make(chan int)
	var wg sync.WaitGroup

	// Launch workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, dataChannel)
	}

	// Ctrl+C signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	// Add data to chan
	go func() {
		defer close(dataChannel)
		for i := 1; ; i++ {
			select {
			case dataChannel <- i:
				time.Sleep(time.Second) // Pause between adding data to chan
			case <-sigCh:
				fmt.Println("Received Ctrl+C signal. Shutting down...")
				return
			}
		}
	}()

	// Waiting for goroutines
	wg.Wait()
	fmt.Println("All workers have shut down. Exiting...")
}

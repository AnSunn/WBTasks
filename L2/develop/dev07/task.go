package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== combine channel ===
Реализовать функцию, которая будет объединять один или более done-каналов в single-канал,
если один из его составляющих каналов закроется.
Очевидным вариантом решения могло бы стать выражение с использованием select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var combineChannels func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“done after %v”, time.Since(start))
*/
// the function takes a variable number of arguments, that's why '...'
func combineChannels(channels ...<-chan interface{}) <-chan interface{} {
	wg := new(sync.WaitGroup)
	out := make(chan interface{})
	//iterate by channels and create new work thread go func() and count it (wg.Add())
	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan interface{}) {
			defer wg.Done()
			//listen to ch and add data to combined 'out' channel. After that close work thread (wg.Done())
			for val := range ch {
				out <- val
			}
		}(ch)
	}
	//wait for all work threads and close output channel
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})

	go func() {
		defer close(c)
		time.Sleep(after)
	}()

	return c
}
func main() {
	start := time.Now()
	<-combineChannels(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
	)
	fmt.Printf("Done after %v\n", time.Since(start))
}

Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```
Ответ
```
0
1
2
3
4
5
6
7
8
9
fatal error: all goroutines are asleep - deadlock!

Цикл for-range читает пока канал не будет закрыт. Здесь канал не закрывается, поэтому после вывода всех значений будет deadlock -
заблокировалась горутина, потому что ожидается закрытие канала


```

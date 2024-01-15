Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
1
3
5
6
7
8
2
4
0
0
...            
Вывод может быть в любом порядке, так как не гарантируется порядок записи в merge канал.
После этого будет происходить чтение из канала, в котором оказываются одни нули, потому что каналы закрылись.
Если канал закрыт, чтение из него всегда будет возвращать нулевое значение для типа канала. 
В данном случае, это будет 0 для int.
Избежать такого поведения можно добавив ok флаги к чтению из закрытых каналов.
```
    data, ok := <-ch
        if !ok {
            fmt.Println("Channel closed")
            break
        }
```
Также необходимо
добавить закрытие merge канала, чтобы программа не схватила deadlock, потому что
range c продолжает читать из канала c, пока он не будет закрыт. У нас канал не закрывается и поэтому бесконечно выводятся нулевые значения, 
а range ожидает закрытия.
```

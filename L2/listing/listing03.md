Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false

Foo() возвращает error. Error - это интерфейсный тип. Интерфейсный тип равен nil только когда его тип и значение равны nil, 
а не когда только значение равно nil. Здесь тип *os.PathError - значит интерфейс уже не nil, поэтому false

```

Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```
Ответ
```
error

так как функция test() возвращает указатель на структуру customError и результат кладется в err, который является интерфейсным типом (error является интерфейсным типом)
интерфейс равен nil только если тип и значение равны nil. А тут тип *customerError, значит попадаем в ветку != nil


```

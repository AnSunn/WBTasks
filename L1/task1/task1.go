package task1

import "fmt"

type human struct {
	name string
	age  int
}

type action struct {
	human
}

func (h *human) greeting() {
	fmt.Printf("I am %s. I am %d years old.\n", h.name, h.age)
}

func (a *action) walk() {
	fmt.Println(a.human.name, "is walking.")
}

func Launch() {
	person := action{
		human: human{
			name: "Alex",
			age:  45,
		},
	}

	person.greeting()
	person.walk()
}

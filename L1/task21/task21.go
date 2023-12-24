package task21

import "fmt"

// Target interface
type Target interface {
	Request() string
}

// Adaptee struct
type Adaptee struct {
}

// func for adaptee struct
func (a *Adaptee) SpecificRequest() string {
	return "Specific request"
}

// Adapter
type Adapter struct {
	Adaptee *Adaptee
}

// func of target interface
func (a *Adapter) Request() string {
	return a.Adaptee.SpecificRequest()
}

func Launch() {

	adaptee := &Adaptee{}
	adapter := &Adapter{Adaptee: adaptee}

	var target Target
	target = adapter
	result := target.Request()

	fmt.Println(result)
}

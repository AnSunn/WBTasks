package pattern

import "fmt"

/*
Implementation of "strategy" pattern - behavioral pattern (https://en.wikipedia.org/wiki/Strategy_pattern)

Applicability of strategy pattern:
1. Variations in Behavior
2. Multiple Algorithms
3. Isolate details of realization
Hide details inside strategy classes

Pros:
1. Hot-swappable algorithms on the fly
2. Isolates algorithm code and data from other classes
3. Moving away from inheritance to delegation
4.Open-Closed Principle
You want to add new algorithms without modifying existing client code, following the Open/Closed Principle.

Cons:
1. Complicates the program through additional classes
2. The client has to know the difference between the strategies in order to choose the appropriate one.

*/

// strategy interface
type StrategyDeliveryType interface {
	Deliver()
}

// concrete delivery type
type Courier struct {
	homeAddress string
}

func (c *Courier) Deliver() {
	fmt.Println("The courier arrives at selected time")
}

type Self struct {
	deliveryPointName string
}

func (s *Self) Deliver() {
	fmt.Println("The order arrives at selected delivery point")
}

// context refers to one of the strategies
type Context struct {
	strategy StrategyDeliveryType
}

func (c *Context) SetContext(s StrategyDeliveryType) {
	c.strategy = s
}

func (c *Context) Deliver() {
	c.strategy.Deliver()
}

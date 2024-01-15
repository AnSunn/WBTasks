package pattern

import (
	"fmt"
	"math"
)

/*
Implementation of "visitor" pattern - behavioral pattern (https://en.wikipedia.org/wiki/Visitor_pattern)

Applicability of visitor pattern:
1. Structure with Multiple Element Types
Use the Visitor pattern when you have a complex object structure with multiple types of elements, and you want to define operations
that can be applied to these elements without modifying their classes.
2. Operations That Vary and Extend
Apply the pattern when you anticipate adding new operations to the elements and want to keep the operations separate
from the element classes.
3. Avoiding Modification of Element Classes
Use the Visitor pattern when you want to avoid modifying the classes of the elements each time a new operation is added,
promoting open-closed principle.

Pros:
1. Extensibility
It allows the addition of new operations without modifying existing element classes, promoting code extensibility.
2. Maintainability
The pattern makes it easier to maintain and modify the code as new operations or changes to existing operations
can be localized within the visitor classes.

Cons:
1. Increased Number of Classes
Introducing the Visitor pattern may lead to a larger number of classes, especially if the object structure has many
different types of elements.
2. Complexity for Simple Structures
For simple object structures, the Visitor pattern might introduce unnecessary complexity, and alternative solutions
may be more straightforward.

Real examples of visitor pattern:
1. Document Processing
2. GUI Components
In graphical user interface libraries, you may have a hierarchy of UI elements. A visitor pattern can be applied to
implement different behaviors or actions that can be performed on these UI elements without modifying their classes.
3. Financial Calculations
In a financial application where you have different types of financial instruments, a visitor pattern could be
used to perform various calculations or reporting operations on these instruments without modifying their classes.
*/

const pi = 3.14

type shape interface {
	accept(visitor) float64
}

// Circle - concrete shape
type Circle struct {
	radius int
}

func NewCircle(r int) *Circle {
	return &Circle{
		radius: r,
	}
}

func (c *Circle) Accept(v visitor) float64 {
	return v.visitCircle(c)
}

// square - concrete shape
type Square struct {
	a int
	b int
}

func NewSquare(a, b int) *Square {
	return &Square{
		a: a,
		b: b,
	}
}

func (s *Square) Accept(v visitor) float64 {
	return v.visitSquare(s)
}

// visitor interface
type visitor interface {
	visitCircle(*Circle) float64
	visitSquare(square *Square) float64
}

// concrete visitor area
type AreaCalculator struct {
	area float64
}

func (a *AreaCalculator) visitCircle(c *Circle) float64 {
	fmt.Println("Calculating circle area")
	return math.Pow(float64(c.radius), 2) * pi
}

func (a *AreaCalculator) visitSquare(s *Square) float64 {
	fmt.Println("Calculating square area")
	return float64(s.a * s.b)
}

// concrete visitor perimeter
type Perimeter struct {
	perimeter float64
}

func (p *Perimeter) visitCircle(c *Circle) float64 {
	fmt.Println("Calculating circle perimeter")
	return 2 * pi * float64(c.radius)
}

func (p *Perimeter) visitSquare(s *Square) float64 {
	fmt.Println("Calculating square perimeter")
	return float64(2*s.a + 2*s.b)
}

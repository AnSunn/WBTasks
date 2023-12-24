package task24

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// NewPoint Constructor
func NewPoint(x, y float64) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

// Distance
func (p1 *Point) DistanceTo(p2 *Point) float64 {
	distanceX := p1.X - p2.X
	distanceY := p1.Y - p2.Y
	return math.Sqrt(distanceX*distanceX + distanceY*distanceY)
}
func Launch() {
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	distance := point1.DistanceTo(point2)
	fmt.Printf("The distance is : %.2f\n", distance)
}

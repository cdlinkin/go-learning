/*
exercise:

	Shape with Area() and Perimeter() methods; types Circle, Rect, Triangle;
	function for the total area []Shape; sort.Slice by area

exercise:

	Implement String() for your type and make sure that fmt.Println uses it.
*/
package stage2

import (
	"fmt"
	"math"
	"sort"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	R float64
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.R, 2)
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

func (c *Circle) String() string {
	return fmt.Sprintf("Circle: Area: %.2f.\n Perimeter: %.2f.", c.Area(), c.Perimeter())
}

type Rect struct {
	A float64
	B float64
}

func (r *Rect) Area() float64 {
	return r.A * r.B
}

func (r *Rect) Perimeter() float64 {
	return 2 * (r.A + r.B)
}

type Triangle struct {
	A float64
	B float64
	C float64
}

func (t *Triangle) Area() float64 {
	p := (t.A + t.B + t.C) / 2
	return math.Sqrt((p * (p - t.A) * (p - t.B) * (p - t.C)))
}

func (t *Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

func sumShapes(shapes []Shape) float64 {
	var sum float64

	for _, v := range shapes {
		sum += v.Area()
	}

	sort.Slice(shapes,
		func(i, j int) bool {
			return shapes[i].Area() < shapes[j].Area()
		},
	)
	return sum
}

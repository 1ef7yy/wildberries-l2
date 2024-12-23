package main

import "fmt"

type Shape interface {
	GetType() string
	Accept(Visitor)
}

type Square struct {
	side float64
}

func (s *Square) Accept(v Visitor) {
	v.VisitSquare(s)
}

func (s *Square) GetType() string {
	return "Square"
}

type Circle struct {
	radius float64
}

func (c *Circle) Accept(v Visitor) {
	v.VisitCircle(c)
}

func (c *Circle) GetType() string {
	return "Circle"
}

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitRectangle(r)
}

func (r *Rectangle) GetType() string {
	return "Rectangle"
}

type Visitor interface {
	VisitSquare(*Square)
	VisitCircle(*Circle)
	VisitRectangle(*Rectangle)
}

type AreaCalculator struct {
	area float64
}

func (a *AreaCalculator) VisitSquare(s *Square) {
	a.area = s.side * s.side
	fmt.Printf("Square area: %f\n", a.area)
}

func (a *AreaCalculator) VisitCircle(c *Circle) {
	a.area = 3.14 * c.radius * c.radius
	fmt.Printf("Circle area: %f\n", a.area)
}

func (a *AreaCalculator) VisitRectangle(r *Rectangle) {
	a.area = r.width * r.height
	fmt.Printf("Rectangle area: %f\n", a.area)
}

// func main() {
// 	square := &Square{side: 5}
// 	circle := &Circle{radius: 3}
// 	rectangle := &Rectangle{width: 4, height: 6}
// 	visitor := &AreaCalculator{}
// 	square.Accept(visitor)
// 	circle.Accept(visitor)
// 	rectangle.Accept(visitor)
// }

// Output:
// Square area: 25.000000
// Circle area: 28.260000
// Rectangle area: 24.000000

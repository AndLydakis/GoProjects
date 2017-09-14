package main

import (
	"fmt";
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type MultiShape struct {
	shapes []Shape
}

func totalArea(ms MultiShape) float64 {
	totalarea := 0.0
	for _, v := range ms.shapes {
		totalarea += v.area()
	}
	return totalarea
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Circle struct {
	x, y, r float64
}

type Rectange struct {
	x1, y1, x2, y2 float64
}

func (r *Rectange) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectange) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l*2 + w*2
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return math.Pi * c.r * 2
}

func (m *MultiShape) area() float64 {
	totalarea := 0.0
	for _, v := range m.shapes {
		totalarea += v.area()
	}
	return totalarea
}

func (m *MultiShape) perimeter() float64 {
	perimeter := 0.0
	for _, v := range m.shapes {
		perimeter += v.perimeter()
	}
	return perimeter
}
func main() {
	multiShape := MultiShape{
		shapes: []Shape{
			&Circle{0, 0, 5},
			&Rectange{0, 0, 10, 10},
		},
	}
	fmt.Println(multiShape.area())
	fmt.Println(multiShape.perimeter())
}

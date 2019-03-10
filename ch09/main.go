package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)

	return l * w
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is ", p.Name)
}

type Android struct {
	Person
	Model string
}

type Shape interface {
	area() float64
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64

	for _, s := range m.shapes {
		area += s.area()
	}

	return area
}

func totalArea(shapes ...Shape) float64 {
	var area float64

	for _, s := range shapes {
		area += s.area()
	}

	return area
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1

	return math.Sqrt(a*a + b*b)
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}
	a := new(Android)

	fmt.Printf("%9.2f\n", c.area())
	fmt.Printf("%9.2f\n", r.area())
	a.Person.Talk()
	a.Talk()

	fmt.Printf("%9.2f\n", totalArea(&c, &r))
}

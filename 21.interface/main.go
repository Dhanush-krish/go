package main

import "fmt"

//stores methods without implementations

type Shape interface {
	Area() float32
}

func Calculate(s Shape) {
	fmt.Println("Area", s.Area())
}

// square class
type Square struct {
	Radius float32
}

// square method
func (s Square) Area() float32 {
	return s.Radius * s.Radius
}

// rectangle class
type Rectangle struct {
	Length float32
	Breath float32
}

// rectangle method
func (r Rectangle) Area() float32 {
	return r.Length * r.Breath
}

func main() {
	fmt.Println("Interface is go")
	// fmt.Println("area of square", Square{10.0}.Area())
	// fmt.Println("area of rectangle", Rectangle{5.0, 6.0}.Area())
	sw := Square{10.0}
	Calculate(sw)

}

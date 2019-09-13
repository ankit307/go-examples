package main

import (
	"fmt"
)

type rect struct {
	height, width float64
}

func (r *rect) area() float64 {
	return r.height * r.width
}

func (r *rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func main() {
	rect := rect{height: 20, width: 10}
	fmt.Printf("Area of react: %f\n", rect.area())
	fmt.Printf("Perimeter is :%f\n", rect.perim())

}

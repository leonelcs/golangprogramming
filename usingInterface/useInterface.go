package main

import (
	"fmt"
	"math"

	"github.com/leonelcs/golangprogramming/myinterface"
)

type square struct {
	X float64
}

type circle struct {
	R float64
}

type retangle struct {
	H float64
	W float64
}

func (s square) Area() float64 {
	return s.X * s.X
}

func (s square) Perimeter() float64 {
	return 4 * s.X
}

func (c circle) Area() float64 {
	return c.R * c.R * math.Pi
}

func (c circle) Perimeter() float64 {
	return 2 * c.R * math.Pi
}

func (r retangle) Area() float64 {
	return r.H * r.W
}

func (r retangle) Perimeter() float64 {
	return (2 * r.H) + (2 * r.W)
}

//Calculate Area and Perimeter of Shape
func Calculate(x myinterface.Shape) {
	switch v := x.(type) {
	case square:
		fmt.Println("This is a square!")
	case circle:
		fmt.Println("This is a circle!", v)
	case retangle:
		fmt.Println("This a retangle!")
	default:
		fmt.Printf("Unknown type %T!\n", v)
	}

	fmt.Println(x.Area())
	fmt.Println(x.Perimeter())
}

func main() {

	x := square{X: 10}
	fmt.Println("Perimeter:", x.Perimeter())
	Calculate(x)
	y := circle{R: 5}
	Calculate(y)

}

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	// 5
	fmt.Println(AbsFunc(v))
	// 5

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	// 5
	fmt.Println(AbsFunc(*p))
	// 5
}

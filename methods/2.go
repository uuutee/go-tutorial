package main

import (
	"math"
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X)
}

func main()  {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
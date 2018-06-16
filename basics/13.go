package main

import (
	"math"
	"fmt"
)

func main() {
	var x, y = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

// 3 4 5
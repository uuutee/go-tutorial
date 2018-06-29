package main

import (
	"math"
	"fmt"
)

type Vertex struct {
	X, Y float64
}

// レシーバ引数を設定することで、structに関数を定義する
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

// 5
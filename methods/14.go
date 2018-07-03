package main

import "fmt"

func main() {
	var i interface{}
	describe(i)
	// (<nil>, <nil>)

	i = 42
	describe(i)
	// (42, int)

	i = "hello"
	describe(i)
	// (hello, string)
}

// 空のインターフェースにはどんな型でも入れられる
func describe(i interface{})  {
	fmt.Printf("(%v, %T)\n", i, i)
}

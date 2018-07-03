package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I

	describe(i)
	// (<nil>, <nil>)

	// i.M()
	// nilのインターフェース呼び出しはエラーになる
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x20 pc=0x10951b8]
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

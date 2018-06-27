package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**5d = %d\n", i, v)
	}
}

// 2**5d = 0
// %!(EXTRA int=1)2**5d = 1
// %!(EXTRA int=2)2**5d = 2
// %!(EXTRA int=4)2**5d = 3
// %!(EXTRA int=8)2**5d = 4
// %!(EXTRA int=16)2**5d = 5
// %!(EXTRA int=32)2**5d = 6
// %!(EXTRA int=64)2**5d = 7
// %!(EXTRA int=128)

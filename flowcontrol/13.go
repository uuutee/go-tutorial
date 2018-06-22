package main

import "fmt"

func main() {
	fmt.Println("counting")

	// LIFO で実行される
	for i :=0; i < 10; i++ {
		defer println(i)
	}

	fmt.Println("done")
}

// counting
// done
// 9
// 8
// 7
// 6
// 5
// 4
// 3
// 2
// 1
// 0
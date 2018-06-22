package main

import "fmt"

func main() {
	defer fmt.Println("world") // main() の終了時に実行される
	fmt.Println("hello")
}

// hello
// world
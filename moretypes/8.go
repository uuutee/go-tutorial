package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names) // [John Paul George Ringo]

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) // [John Paul] [Paul George]

	// スライスを変更するとスライス元の配列も変更される
	b[0] = "xxx"
	fmt.Println(a, b) // [John xxx] [xxx George]
	fmt.Println(names) // [John xxx George Ringo]
}

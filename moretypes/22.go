package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The Value:", m["Action"]) // The Value: 0

	m["Answer"] = 48
	fmt.Println("The Value:", m["Answer"]) // The Value: 48

	delete(m, "Answer")
	fmt.Println("The Value:", m["Answer"]) // The Value: 0

	// key exists的なやつ
	v, ok := m["Answer"]
	fmt.Println("The Value:", v, "Present?", ok) // The Value: 0 Present? false
}

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
} 

func (t *T) M()  {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	// (<nil>, %!t(*main.T=<nil>))
	i.M()
	// <nil>

	i = &T{"hello"}
	describe(i)
	// (&{hello}, &{%!t(string=hello)})
	i.M()
	// hello
}

func describe(i I)  {
	fmt.Printf("(%v, %t)\n", i, i)
}

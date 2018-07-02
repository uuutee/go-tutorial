package main

import "fmt"

type I interface {
	M()
}
type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
// interfaceの実装は、メソッドを実装するだけ。他言語のimplements のようなキーワード指定はいらない。
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}

// hello

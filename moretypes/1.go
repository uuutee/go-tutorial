package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	p := &i // iへのポインタ
	fmt.Println(i)

	p = &j // jへのポインタ
	*p = *p / 37 // jをポインタp経由で割る
	fmt.Println(j)
}

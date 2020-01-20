package main

import (
	"fmt"
)

func main() {

	g := func(xi []int) int {
		return len(xi)
	}
	x := foo(g, []int{1, 2, 3, 4, 5, 6})
	fmt.Println(x)
}

func foo(f func(xi []int) int, ii []int) int {
	n := f(ii)
	return n
}

package main

import (
	"fmt"
)

func main() {

	x := foo()
	fmt.Println(x())
	fmt.Println(x())
	y := foo()
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())


}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

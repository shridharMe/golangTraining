package main

import "fmt"

func main() {

	fmt.Println("I am learning control flow in goLang")

	foo()

	fmt.Println("some loop and conditional example")

	for i := 0; i < 100; i++ { // loop
		if i%2 == 0 { // conditional
			fmt.Println(i)
		}
	}
	bar()

}

func foo() {
	fmt.Println("I am, in foo")
}

func bar() {
	fmt.Println("No I understood, How control flow works in  goLang")
}

//Control flow
// (1) sequence
// (2) loop: iterative
// (3) conditional

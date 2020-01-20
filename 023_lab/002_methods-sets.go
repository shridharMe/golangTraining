package main

import (
	"fmt" 
)

type person struct {
	first string
	last  string
}

func (a *person) speak() {
	fmt.Println("speaking hello")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		first: "james",
		last:  "bond",
	}

	saySomething(&p)

	p.speak()
}

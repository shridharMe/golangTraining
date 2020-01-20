package main

import (
	"fmt"
)
 type person struct {
	 first string
	 last string
	 age int
 }

 func (per person) speak(){
    fmt.Println(per.first, per.last, per.age)
 } 

func main () {	 
	p := person {
		 first: "shridhar",
		 last: "patil",
		 age: 32,
	}
	p.speak()
}
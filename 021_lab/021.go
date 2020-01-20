package main

import (
	"fmt"
)

func exercise1(){
	x := 42
	fmt.Println(x)
	fmt.Println(&x)
}
type person struct {
	first string
	last string
	age int
}

func changeMe(p *person){
		(*p).first="james"
}

func exercise2(){
	p := person{
		first: "Dr.",
		last: "no",
		age : 50,
	}

	fmt.Println(p.first)
   	changeMe(&p)
  	fmt.Println(p.first)

}
func main(){
	//exercise1()
	exercise2()
}
package main

import (
	"fmt"
)

var x func() 

func boo() func (){
	return func (){
		fmt.Println("I am anonymous return")
	}
} 
 
func main(){
	 x = boo()

	 x()
}
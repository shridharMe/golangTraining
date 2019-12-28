package main

import (
	"fmt"
)
var a int
type hotdog int
var b hotdog
func main(){
	a = 20
    b= 10
 
	fmt.Println(a)
	fmt.Printf("%T\n" , a )

	fmt.Println(b)
	fmt.Printf("%T\n" , b )

	convert()
}

func convert(){
	a = int(b)
	 
	fmt.Println(a)
	fmt.Printf("%T\n" , a )
}
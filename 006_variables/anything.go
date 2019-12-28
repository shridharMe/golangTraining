package main

import (
	"fmt"
)

var y = 42
var z int
var s string = "hello india"
var a string = ` Good morning,

"hello india"`

func main (){
 fmt.Println( y)
	foo()
	z = 10
	fmt.Println( z)
	fmt.Println( s)
	fmt.Println( a)

	typechecking()
}

func foo(){
 fmt.Println( z)
}

func typechecking(){
 fmt.Printf("%T\n",y)
 fmt.Printf("%T\n",s)
 fmt.Printf("%T\n",z)

}
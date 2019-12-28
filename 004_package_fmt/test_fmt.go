package main


import (
	"fmt"
)

func main(){

n,_ := fmt.Println("hello world",42, true)
fmt.Println(n)

//output : 
//hello india 42 true 
//20

n, err := fmt.Println("hello india",42, true)
fmt.Println(n,err)
//output :

//hello india 42 true 
//20 <nil>


//n, ed := fmt.Println("hello error",42, true)
//fmt.Println(n)
//output : ed declared and not used
}
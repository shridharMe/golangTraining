package main

import (
	"fmt"
)

const (
	_ = iota
	a = 2019 + iota
	b
	c
	d

)

func main(){
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Printf("%d\n%b\n%x",42,42,42)
}
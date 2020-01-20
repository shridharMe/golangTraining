package main

import (
	"fmt"
)

func foo () int{
  return 43
}

func bar () (int, string){
  return 43, "shridhar"
}



func main () {
	fmt.Println(foo())
	fmt.Println(bar())
}
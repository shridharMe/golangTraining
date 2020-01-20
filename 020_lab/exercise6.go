package main

import (
  "fmt"
   
)

func main(){
	func (){
		fmt.Println("I am anonymous")
	}()

	x :=  func (){
		fmt.Println("I am anonymous 2")
	}

	x()
}
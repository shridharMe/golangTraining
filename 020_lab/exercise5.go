package main

import (
  "fmt"
   
)
 type SQUARE struct {
	  length int	 
 }

 type CIRCLE struct {
	  radius int  
 }

func (s SQUARE) area()int{
    return s.length*s.length
 } 

func (c CIRCLE) area()int{
     return  c.radius*c.radius
 } 

type shape interface {
	area() int
}

func info(s shape) {
	fmt.Println(s.area())
}

func main () {	 
	sq := SQUARE{
		length: 15,
	}
 
 cr := CIRCLE{
	 radius: 15,
 }
info(cr)
	info(sq)
}
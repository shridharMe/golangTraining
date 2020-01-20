package main

import (
	"fmt"
)

func exercise5() {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	v, ok := <-c
	fmt.Println(v, ok)
	close(c)
	v, ok = <-c
	fmt.Println(v, ok)
}
func exercise6() {
	c := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			c <- i
		}
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}
func exercise7() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 1; j <= 10; j++ {
				c <- j
			}

		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

}

func main() {
	exercise7()
}

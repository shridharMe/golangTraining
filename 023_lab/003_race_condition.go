package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {
	var wg sync.WaitGroup
	incrementer := 0
	gs := 100
	wg.Add(gs)
	
	for i:=0; i<gs;i++{
		go func(){
			
			v := incrementer // read the increament value
			runtime.Gosched()
			v++
			incrementer=v
			fmt.Println(incrementer)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}


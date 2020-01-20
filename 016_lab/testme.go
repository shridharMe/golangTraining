package main

import (
	"fmt"
)

func main() {
	/*fmt.Println("Excercixe #2")
	exercise2()
	fmt.Println("Excercixe #3")
	exercise3()
	fmt.Println("Excercixe #4")
	exercise4()
	fmt.Println("Excercixe #5")
	exercise5()*/
	fmt.Println("Excercixe #6")
	exercise6()

}

func exercise2() {
	for i := 65; i < 91; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t%#U\n", i)
		}

	}
}

func exercise3() {
	for i := 1976; i <= 2019; i++ {
		fmt.Println(i)
	}

}

func exercise4() {
	x :=1976
	for  {
		if  x <= 2019 {
			fmt.Println(x)
		}
		 x++
		
	}

}

func exercise5() {
	x :=10
	for  {
		if  x <= 100 {
			fmt.Println(x%4)
		}
		 x++
		
	}

}

func exercise6() {
	favSport:="surfing"
	switch favSport  {
	case "surfing":
			fmt.Println("I am in")
		default :
			fmt.Println("default value")
	}

}

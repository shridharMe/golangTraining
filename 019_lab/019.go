package main

import (
	"fmt"
)

type person struct {
	firstname    string
	lastname     string
	icecreamflvr string
}

func exercise1() {
	p1 := person{
		firstname:    "shridhar",
		lastname:     "patil",
		icecreamflvr: "vanila",
	}

	p2 := person{
		firstname:    "james",
		lastname:     "bond",
		icecreamflvr: "choclate",
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func exercise2() {
	p1 := person{
		firstname:    "shridhar",
		lastname:     "patil",
		icecreamflvr: "vanila",
	}

	p2 := person{
		firstname:    "james",
		lastname:     "bond",
		icecreamflvr: "choclate",
	}

	m := map[string][]string{
		p1.lastname: []string{p1.firstname, p1.icecreamflvr},
		p2.lastname: []string{p2.firstname, p2.icecreamflvr},
	}
	for k, v := range m {
		fmt.Println(k)
		for k2, v2 := range v {
			fmt.Printf("\t\t%v\t\t%v\n", k2, v2)
		}
	}

}

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func exercise3() {

	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Printf("\t\t%v", t)
	fmt.Printf("\t\t%v", s)
}

func exercise4() {
	t := struct {
		vehicle
		luxury bool
	}{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		luxury: true,
	}

	fmt.Printf("\t\t%v", t)
}

func main() {
	exercise4()

}

package main

import (
	"fmt"
)
//array
func exercise1() {
	x := [5]int{1, 2, 3, 4, 5}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}

//slice
func exercise2() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}

//SLICING on slice
func exercise3() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(x[0:5])
	fmt.Println(x[6:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}

//appending to Slice
func exercise4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	x = append(x, 53, 54, 55)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

func exercise5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := x[0:3]
	y = append(y, x[6:]...)
	fmt.Println(y)

}

func exercise6() {
	x := []string{
		` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
	}
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
//multi dimensional slice
func exercise7() {
	x1 := []string{"James", "Bond", "Shaken, not stirred"}
	x2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	x := [][]string{x1, x2}
	fmt.Println(x)

	for i, v := range x {
		for j, vs := range v {
			fmt.Println(i, j, vs)
		}

	}

}
//map
func exercise8() {
	m := map[string][]string{
	`bond_james`: []string{`Shaken, not stirred`, `Martinis`, `Women`},
	`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
	`no_dr`:[]string{ `Being evil`, `Ice cream`, `Sunsets`},
	}

	//fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)
		for k2, v2 := range v {
			fmt.Printf("\t\t%v\t\t%v\n",k2,v2)
		}

	}
}
// add to map
func exercise9() {
	m := map[string][]string{
	`bond_james`: []string{`Shaken, not stirred`, `Martinis`, `Women`},
	`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
	`no_dr`:[]string{ `Being evil`, `Ice cream`, `Sunsets`},
	}

	m[`patil_shridhar`] = []string{`eating`,`reading`,`programming`}

	for k, v := range m {
		fmt.Println(k)
		for k2, v2 := range v {
			fmt.Printf("\t\t%v\t\t%v\n",k2,v2)
		}

	
	}

}


// delete to map
func exercise10() {
	m := map[string][]string{
	`bond_james`: []string{`Shaken, not stirred`, `Martinis`, `Women`},
	`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
	`no_dr`:[]string{ `Being evil`, `Ice cream`, `Sunsets`},
	}

	m[`patil_shridhar`] = []string{`eating`,`reading`,`programming`}

	 
	delete(m, `no_dr`)

	
	for k, v := range m {
		fmt.Println(k)
		for k2, v2 := range v {
			fmt.Printf("\t\t%v\t\t%v\n",k2,v2)
		}

	
	}
}

func main() {
	//exercise1()
	//exercise2()
	//exercise3()
	//exercise4()
	//exercise5()
	//exercise6()
	//exercise7()
	//exercise8()
	//exercise9()
	exercise10()
}
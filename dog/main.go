
// package dog  have an exported func “Years” which takes human years and turns them into dog years (1 human year = 7 dog years)
package dog

// Years converts human years to dog years.
func Years(humanYears int)(dogYears int){
	return humanYears*7
}

// YearsTwo converts human years to dog years.
func YearsTwo(humanYears int) int {
	dogYears := 0
	for i := 0; i < humanYears; i++ {
		dogYears += 7
	}
	return dogYears
}
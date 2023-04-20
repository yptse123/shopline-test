// Given that there is an array containing numbers from 1 to 100.
// When the number is multiple of 3, print "bug"
// When the number is multiple of 5, print "fix"
// When the number is multiple of 3 and 5, print "bugfix"

package lib

import "fmt"

func Q1() {
	fmt.Println("----------------- Start of Question 1 -----------------")
	// Create an array of numbers from 1 to 100
	numbers := make([]int, 100)
	for i := 0; i < 100; i++ {
		numbers[i] = i + 1
	}

	// Iterate over the array
	for _, number := range numbers {
		// Check if the number is a multiple of 3 and/or 5
		if number%3 == 0 && number%5 == 0 {
			fmt.Println("bugfix")
		} else if number%3 == 0 {
			fmt.Println("bug")
		} else if number%5 == 0 {
			fmt.Println("fix")
		}
	}
	fmt.Println("----------------- End of Question 1 ----------------- \n \n ")
}

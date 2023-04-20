// Given two arrays, [1,2,3,4,5] and [2,3,1,0,5]
// find which number(s) is/are not present in the second array

package lib

import "fmt"

func Q2() {
	fmt.Println("----------------- Start of Question 2 -----------------")
	// Create two arrays
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{2, 3, 1, 0, 5}

	// Create map to store the values from the 2nd array
	present := make(map[int]bool)
	for _, num := range arr2 {
		present[num] = true
	}

	// Iterate over the first array and check if value is present
	notPresent := []int{}
	for _, num := range arr1 {
		if !present[num] {
			notPresent = append(notPresent, num)
		}
	}

	// Print the values
	fmt.Println("Not present:", notPresent)

	fmt.Println("----------------- End of Question 2 -----------------  \n \n ")
}

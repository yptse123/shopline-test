// Given two arrays, [1,2,3,4,5] and [2,3,1,0,5]
// find which number(s) is/are common in both array

package lib

import "fmt"

func Q3() {
	fmt.Println("----------------- Start of Question 3 -----------------")
	// Create the two arrays
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{2, 3, 1, 0, 5}

	// Create a map to store the values from the second array
	present := make(map[int]bool)
	for _, num := range arr2 {
		present[num] = true
	}

	// Iterate over the first array and check if each value is present in the map
	common := []int{}
	for _, num := range arr1 {
		if present[num] {
			common = append(common, num)
		}
	}

	// Print the values
	fmt.Println("Common values:", common)

	fmt.Println("----------------- End of Question 3 -----------------  \n \n ")
}

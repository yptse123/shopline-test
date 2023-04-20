// Given two arrays, [1,2,3,4,5] and [2,3,1,0,5]
// merge these two arrays and unique to display

package lib

import "fmt"

func Q4() {
	fmt.Println("----------------- Start of Question 4 -----------------")
	// Define two arrays
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{2, 3, 1, 0, 5}

	// Initialize a map to store unique elements
	unique := make(map[int]bool)

	// Iterate over the arr1 and add them to the unique map
	for _, num := range arr1 {
		unique[num] = true
	}

	// Iterate over the arr2 and add any new unique elements to the map
	for _, num := range arr2 {
		if _, ok := unique[num]; !ok {
			arr1 = append(arr1, num)
			unique[num] = true
		}
	}

	// Sort the merged slice in ascending order
	for i := 0; i < len(arr1); i++ {
		for j := i + 1; j < len(arr1); j++ {
			if arr1[j] < arr1[i] {
				arr1[i], arr1[j] = arr1[j], arr1[i]
			}
		}
	}

	// Print the values
	fmt.Println("Unique values:", arr1)

	fmt.Println("----------------- End of Question 4 -----------------  \n \n ")
}

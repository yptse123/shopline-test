// You have to design a programme/ function to achieve below objectives:
// How do you find the closest integer 17 in [30, 1, 5, 16, 19, 21, 2, 55]?

package lib

import (
	"fmt"
	"math"
)

func Q5() {
	fmt.Println("----------------- Start of Question 5 -----------------")
	// Create the array
	arr := []int{30, 1, 5, 16, 19, 21, 2, 55}

	// Initialize the closest element and difference to the first element in the array
	closest := arr[0]
	diff := math.Abs(float64(closest - 17))

	// Iterate over the rest of the elements in the array
	for _, num := range arr[1:] {
		currDiff := math.Abs(float64(num - 17))
		if currDiff < diff {
			closest = num
			diff = currDiff
		}
	}

	// Print the closest element
	fmt.Println("Closest element: ", closest)

	fmt.Println("----------------- End of Question 5 -----------------  \n \n ")
}

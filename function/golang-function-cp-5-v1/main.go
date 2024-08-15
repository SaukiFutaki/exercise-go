package main

import (
	"fmt"
)

func FindMin(nums ...int) int {
	// return 0 // TODO: replace this
	var min int = nums[0]
	for _, num := range nums {
		if num < min {
			// fmt.Println(num)
			min = num
		}
	} 
	return min
}

func FindMax(nums ...int) int {
	// return 0 // TODO: replace this
	var max int = nums[0]
	for _, num := range nums {
		if num > max {
			// fmt.Println(num)
			max = num
		}
	}
	return max
}

func SumMinMax(nums ...int) int {
	// return 0 // TODO: replace this
	min := FindMin(nums...)
	max := FindMax(nums...)
	return min + max
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(SumMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(FindMin( 2, 3, 1, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(FindMax( 2, 3, 1, 4, 5, 6, 7, 8, 9, 10))
}

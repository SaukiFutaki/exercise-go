package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	// return 0 // TODO: replace this
	numbersStr := strconv.Itoa(numbers)
	var biggest int
	var pair int



	
	for i:= 1; i < len(numbersStr); i++ {
		num1, _ := strconv.Atoi(string(numbersStr[i]))
		num2, _ := strconv.Atoi(string(numbersStr[i-1]))

		if num1 + num2 > biggest {
			biggest = num1 + num2
			pairNum, _ := strconv.Atoi(string(numbersStr[i-1]) + string(numbersStr[i]))
			pair = pairNum

		}
	}

	return pair
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(11223344))
}

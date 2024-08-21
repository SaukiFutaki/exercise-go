package main

import "fmt"

func ExchangeCoin(amount int) []int {
	var money = []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}	
	var result []int
	if amount == 0 {
		return []int{}
	}
	for _, coin := range money {
		// fmt.Println( coin)
		for amount >= coin {
			result = append(result, coin)
			amount -= coin
		}
	}
	return result
}


func main() {
	fmt.Println(ExchangeCoin(1752))
}
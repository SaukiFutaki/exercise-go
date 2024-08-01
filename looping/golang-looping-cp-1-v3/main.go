package main

import "fmt"

func CountingNumber(n int) float64 {
	var total float64
	for i := 1.0; i <= float64(n); i = i + 0.5 {
		fmt.Println(i)
		total += i
	}
	return total
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(10))
}

package main

import "fmt"

func ReverseData(arr [5]int) [5]int {
	// return nil // TODO: replace this
	var result [5]int
	for i := 0; i < len(arr); i++ {
		// fmt.Println(arr[len(arr)-1-i])	
		// result[i] = arr[len(arr)-1-i]
		num := arr[len(arr)-1-i]
		reverse := 0
		for num > 0 {
			// TODO : reverse x 10 + num%10
			// TODO : karena apa ?
			// TODO : karena  ingin mengambil digit terakhir dari num
			// TODO : dan memasukkannya ke digit pertama dari reverse
			reverse = reverse*10 + num%10

			// TODO : menghapus digit terakhir dari num
			num /= 10
		}
		result[i] = reverse
	}
	return result
}


func main() {
	
	arr := [5]int{456789, 44332, 2221, 12, 10}
	
	fmt.Println(ReverseData(arr))
}
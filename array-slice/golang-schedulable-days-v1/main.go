package main

import "fmt"

func SchedulableDays(date1 []int, date2 []int) []int {
	// return nil // TODO: replace this
	//input: date1 = [1, 2, 3, 4], date2 = [3, 4, 5]
	//output: [3, 4]

	var result []int
	result = make([]int, 0)
	if len(date1) == 0 || len(date2) == 0 { 
		return []int{}
	}
	// for i := 0; i < len(date1); i++ {
	// 	for j := 0; j < len(date2); j++ {
	// 		if date1[i] == date2[j] {
	// 			result = append(result, date1[i])
	// 		}else if date1[i] != date2[j] {
	// 			return []int{}
	// 		}
	// 	}
	// }
	for _, d1 := range date1 {
		for _, d2 := range date2 {
			if d1 == d2 {
				result = append(result, d1)
				break
			} 
		}
	}
	fmt.Println(result)
	return result
}

func main() {
	date1 := []int{1, 2, 3, 4}
	date2 := []int{3, 4, 5}
	fmt.Println(SchedulableDays(date1, date2))
}

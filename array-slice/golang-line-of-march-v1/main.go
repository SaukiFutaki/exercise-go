package main

import "fmt"

func Sortheight(height []int) []int {
	// return nil // TODO: replace this

	for i := 0; i < len(height); i++ {

		for j := i + 1; j < len(height); j++ {
			fmt.Println("j:", j)
			if height[i] > height[j] {

				t := height[i]
				height[i] = height[j]
				height[j] = t

			}
		}
	}
	return height
}

func main() {

	fruits := [3]string{"apple", "banana", "kiwi"}
	num := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	// d2 := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// }
	// d2 
	// 1 2 3	[0][0] [0][1] [0][2]
	// 4 5 6	[1][0] [1][1] [1][2]
	d3 := [][][]int{
		{{1, 2, 3}, {4, 5, 6}},
		{{7, 8, 9}, {10, 11, 12}},
		{{13, 14, 15}, {16, 17, 18}},
	}
	for _, t:= range d3 {
		fmt.Println(t)
	}
	// d3
	// 1 2 3	[0][0][0] [0][0][1] [0][0][2]
	// 4 5 6	[0][1][0] [0][1][1] [0][1][2]
	// 7 8 9	[1][0][0] [1][0][1] [1][0][2]
	// 10 11 12	[1][1][0] [1][1][1] [1][1][2]
	// fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(fruits[0], num[1][1])
}

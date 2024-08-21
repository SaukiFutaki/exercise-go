package main

import "fmt"

func CountProfit(data [][][2]int) []int {
	// return nil // TODO: replace this


	resutlMap := make(map[int]int)
	for _, cabang := range data {
		fmt.Println("cabang", cabang)
		for bulanke, bulan := range cabang {
			fmt.Println("bulan", bulan)
			// profit = append(profit, bulan[0] - bulan[1])
			income := bulan[0]
			expense := bulan[1]

			fmt.Println("income", income)
			fmt.Println("expense", expense)
			profit := income - expense
			resutlMap[bulanke+1] += profit
		}
	}
	numOfBulan := 0
	for k := range resutlMap {
		if k  > numOfBulan {
			numOfBulan = k
		}
	}
	result := make([]int,numOfBulan)
	for k, v := range resutlMap {
		result[k-1] = v
	}
	return result
}

func main() {
	data := [][][2]int {
	{{1000,5000},{500,200}},
	{{1200,200},{1000,800}},
	{{1000,5000},{500,200}},	
	}
	fmt.Println(CountProfit(data))
}

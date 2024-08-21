package main

import "fmt"

type Product struct {
	Name  string
	Price int
	Tax   int
}

func MoneyChanges(amount int, products []Product) []int {
	uang := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}
	total := 0

	for _, product := range products {
		total += product.Price + product.Tax
	}

	kembalian := amount - total
	if kembalian < 0 {
		return []int{}
	}
	result := []int{}
	for _, uang := range uang {
		for kembalian >= uang {
			result = append(result, uang)
			kembalian -= uang
		}
	}
	fmt.Println(result)
	return result

	// return nil // TODO: replace this
}

func main() {
	//input: 10000, [{Name: "Baju", Price: 5000, Tax: 500}, {Name: "Celana", Price: 3000, Tax: 300}]
	// output: [1000, 200]
	products := []Product{
		{Name: "Baju", Price: 5000, Tax: 500},
		{Name: "Celana", Price: 3000, Tax: 300},

	}
	MoneyChanges(10000, products)
}

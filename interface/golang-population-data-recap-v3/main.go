package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]any {
	result  :=  []map[string]any{}
	if len(data) == 0 {
		fmt.Println("Data is empty")
		return result
	}
	for _, item := range data {

		s := strings.Split(item, ";")
		// fmt.Println(s)
		age, _ := strconv.Atoi(s[1])

		item := map[string]interface{}{
			"name":    s[0],
			"age":     age,
			"address": s[2],
		}
		// fmt.Println(item)

		if len(s) > 3 {
			if height, err := strconv.ParseFloat(s[3], 64); err == nil {
				item["height"] = height
			}
		}
		if len(s) > 4 && s[4] != "" {
			if isMarried, err := strconv.ParseBool(s[4]); err == nil {
				item["isMarried"] = isMarried
			}
		}

		result = append(result, item)
		// splitted := strings.Split(item, ",")
		// name := splitted[0]
		// age := splitted[1]
		// city := splitted[2]
		// // population, _ := strconv.Atoi(splitted[1])
		// // area, _ := strconv.Atoi(splitted[2])
		// // density := float64(population) / float64(area)

		// result = append(result, map[string]interface{}{
		// 	"name": name,
		// 	"age": age,
		// 	"city": city,
		// })
	}
	// fmt.Println(result)
	return result
}

func main() {
	//input: data = ["Budi;23;Jakarta;;", "Joko;30;Bandung;;true", "Susi;25;Bogor;165.42;"]

	data := []string{}

	fmt.Println(PopulationData(data))
}

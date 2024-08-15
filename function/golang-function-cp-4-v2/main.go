package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	// return "" // TODO: replace this
	var result []string
	for _, value := range data {
		// ? fungsi contains digunakan untuk mencari apakah input ada di value
		if strings.Contains(value, input) {
			// result += value + "," 
			result = append(result, value)
		}
	}
	return strings.Join(result, ",")
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))

	
}

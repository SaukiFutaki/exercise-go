package main

import "fmt"

func MapToSlice(mapData map[string]string) [][]string {
	result := make([][]string, 0)
	for key, value := range mapData {
		result = append(result, []string{key, value})
	}
	return result // TODO: replace this
}

func main() {
	mapData := map[string]string{"name": "John", "age": "30", "city": "New York"}
	fmt.Println(MapToSlice(mapData))
}

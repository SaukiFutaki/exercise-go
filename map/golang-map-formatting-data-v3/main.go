package main

import (
	"fmt"
	"strings"
)

// TODO: answer here

func ChangeOutput(data []string) map[string][]string {
	res := make(map[string][]string)
	for _, e := range data {
		p := strings.Split(e, "-")
		header := p[0]
		idx := p[1]
		key := p[2]
		value := p[3]

		if key == "first" {
			res[header] = append(res[header], value)
		} else if key == "last" {
			i := 0
			fmt.Sscanf(idx, "%d", &i)
			res[header][i] += " " + value
		}
	}
	// return nil // TODO: replace this
	return res
}

// bisa digunakan untuk melakukan debug
func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar", "phone-0-first-081234567890", "phone-1-first-081234567891"}
	res := ChangeOutput(data)

	fmt.Println(res)
}

package main

import "fmt"

func CountingLetter(text string) int {
	// unreadable letters = R, S, T, Z
	var total int 
	for i := 0; i < len(text); i++ {
		// if text[i] == 'R' || text[i] == 'S' || text[i] == 'T' || text[i] == 'Z' {
		// 	total++
		// }
		fmt.Println(string(text[i]))
		if (text[i] == 'R' || text[i] == 'r') ||( text[i] == 'S' || text[i] == 's') ||( text[i] == 'T' || text[i] == 't') ||( text[i] == 'Z' || text[i] == 'z')  {
			total++
		}
	}
	
	return total
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Remaja muda yang berbakat"))
}

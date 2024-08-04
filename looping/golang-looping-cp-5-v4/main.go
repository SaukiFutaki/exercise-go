package main

import (
	"fmt"
)

func ReverseWord(str string) string {
	var result string
	var word string
	
	// for i, word := range str {
	// 	fmt.Println(i, string(word))
	// }

	for i := 0; i < len(str); i++ {
		if str[i] == ' '  || i == len(str)-1{
			var reverse string
			for j:= len(word)-1; j >= 0; j-- {
				reverse += string(word[j])
				fmt.Println(reverse)
			}
			result += reverse + " "
		}
		word += string(str[i]) 
	}

	fmt.Println(result)
	return "" // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseWord("Aku Sayang Ibu"))
	// fmt.Println(ReverseWord("A bird fly to the Sky"))
}

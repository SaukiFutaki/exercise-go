package main

import (
	"fmt"
	"unicode"
)

func ReverseWord(str string) string {
	var result string
	var word string
	
	// for i, word := range str {
	// 	fmt.Println(i, string(word))
	// }

	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			word += string(str[i])
			
		}

		if str[i] == ' '  || i == len(str)-1{
			var reverse string
			for j:= len(word)-1; j >= 0; j-- {
				reverse += string(word[j])
				fmt.Println(reverse)
			}

			if unicode.IsUpper(rune(word[0])) == true {
				reverse = string(unicode.ToUpper(rune(reverse[0]))) + reverse[1:]
			}

			if unicode.IsLower(rune(word[len(word)-1])) == true {
				reverse = reverse[:len(reverse)-1] + string(unicode.ToLower(rune(reverse[len(reverse)-1])))
			}
			result += reverse + " "
			word = ""
		}
		// word += string(str[i]) 
	}

	fmt.Println(result)
	// return "" // TODO: replace this
	return result[:len(result)-1]

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseWord("Aku Sayang Ibu"))
	// fmt.Println(ReverseWord("A bird fly to the Sky"))
}

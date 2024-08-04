package main

import (
	"fmt"
	"strings"
	
)

func CountVowelConsonant(str string) (int, int, bool) {
	
	var vowel, consonant int
	var isTrue bool

	for i := 0; i < len(str); i++ {
		lowerChar := strings.ToLower(string(str[i]))
			
		if lowerChar == "a"  || lowerChar == "i" || lowerChar == "u" || lowerChar== "e" || lowerChar == "o" {
			vowel++
		} else if(lowerChar >= "a" && lowerChar <= "z") {
			consonant++
		} else if vowel == 0 && consonant == 0  {
			isTrue = false
		} else if vowel == 0 && consonant > 0 {
			isTrue = true
		}
	}

	fmt.Println(vowel)
	fmt.Println(consonant)
	fmt.Println(isTrue)

	return vowel, consonant, isTrue 
	 // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("bbbbbbbb vvvvvvvv  ddddd sssss  wwww")) 
}

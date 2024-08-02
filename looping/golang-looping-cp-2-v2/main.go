package main

import (
	"fmt"	
)
// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {
	var result string
	for i := len(str) - 1; i >= 0; i-- {
		if (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z') || (str[i] >= '0' && str[i] <= '9') {
			result += string(str[i])
			if i != 0 && ((str[i-1] >= 'a' && str[i-1] <= 'z') || (str[i-1] >= 'A' && str[i-1] <= 'Z') || (str[i-1] >= '0' && str[i-1] <= '9')) {
				result += "_"
			}
		} else if str[i] == ' ' {
			result += " "
		}
	}
	return result
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseString("I am a student."))
}

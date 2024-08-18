package main

import (
	"fmt"
)

func PhoneNumberChecker(number string, result *string) {
	// TODO: answer here

	//0811 - 0815 / 62811 - 62815 => Telkomsel
	//0816 - 0819 / 62816 - 62819 => Indosat
	//0821 - 0823 / 62821 - 62823 => XL
	//0827 - 0829 / 62827 - 62829 => Tri
	//0852 - 0853 / 62852 - 62853 => AS
	//0881 - 0888 / 62881 - 62888 => Smartfren

	// for i, v := range number {
	// 	if i == 0 {
	// 		if v == '0' {
	// 			if number[i+1] == '8' {
	// 				*result = "valid"
	// 				return
	// 			}
	// 		} else if v == '6' {
	// 			if number[i+1] == '2' && number[i+2] == '8' {
	// 				*result = "valid"
	// 				return
	// 			}
	// 		}
	// 	}
	// 	fmt.Println(string(v))
	// }

	
	if len(number) >= 11 && number[:2] == "62" {
		number = "0" + number[2:]
		fmt.Println(number)
	} else if len(number) >= 10 && number[:2] != "08" {
		*result = "invalid"
		return
	} else if len(number) < 10 {
		*result = "invalid"
		return
	}
	firstNumber := number[:4]
	switch firstNumber {
	case "0811", "0812", "0813", "0814", "0815":
		*result = "Telkomsel"
	case "0816", "0817", "0818", "0819":
		*result = "Indosat"
	case "0821", "0822", "0823":
		*result = "XL"
	case "0827", "0828", "0829":
		*result = "Tri"
	case "0852", "0853":
		*result = "AS"
	case "0881", "0882", "0883", "0884", "0885", "0886", "0887", "0888":
		*result = "Smartfren"
	default:
		*result = "invalid"
	}
	// if len(number) > 10 && number[:2] == "62" {
	// 	number = "0" + number[2:]
	// 	fmt.Println(number)
	// 	firstNumber = number[:4]
	// 	if firstNumber == "0811" || firstNumber == "0812" || firstNumber == "0813" || firstNumber == "0814" || firstNumber == "0815" {
	// 		*result = "Telkomsel"
	// 	} else if firstNumber == "0816" || firstNumber == "0817" || firstNumber == "0818" || firstNumber == "0819" {
	// 		*result = "Indosat"
	// 	} else if firstNumber == "0821" || firstNumber == "0822" || firstNumber == "0823" {
	// 		*result = "XL"
	// 	} else if firstNumber == "0827" || firstNumber == "0828" || firstNumber == "0829" {
	// 		*result = "Tri"
	// 	} else if firstNumber == "0852" || firstNumber == "0853" {
	// 		*result = "AS"
	// 	} else if firstNumber == "0881" || firstNumber == "0882" || firstNumber == "0883" || firstNumber == "0884" || firstNumber == "0885" || firstNumber == "0886" || firstNumber == "0887" || firstNumber == "0888" {
	// 		*result = "Smartfren"
	// 	}
	// } else if len(number) == 10 && number[:2] == "08" {
	// 	firstNumber = number[:4]
	// 	if firstNumber == "0811" || firstNumber == "0812" || firstNumber == "0813" || firstNumber == "0814" || firstNumber == "0815" {
	// 		*result = "Telkomsel"
	// 	} else if firstNumber == "0816" || firstNumber == "0817" || firstNumber == "0818" || firstNumber == "0819" {
	// 		*result = "Indosat"
	// 	} else if firstNumber == "0821" || firstNumber == "0822" || firstNumber == "0823" {
	// 		*result = "XL"
	// 	} else if firstNumber == "0827" || firstNumber == "0828" || firstNumber == "0829" {
	// 		*result = "Tri"
	// 	} else if firstNumber == "0852" || firstNumber == "0853" {
	// 		*result = "AS"
	// 	} else if firstNumber == "0881" || firstNumber == "0882" || firstNumber == "0883" || firstNumber == "0884" || firstNumber == "0885" || firstNumber == "0886" || firstNumber == "0887" || firstNumber == "0888" {
	// 		*result = "Smartfren"
	// 	}
	// } else {
	// 	*result = "invalid"
	// 	return

	// }

	fmt.Println(result)
	fmt.Println(firstNumber)

}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "081211111111"
	var result string
	fmt.Println("Nomor telepon: ", number)
	PhoneNumberChecker(number, &result)

}

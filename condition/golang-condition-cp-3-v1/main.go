package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {
	var mean = (math + science + english + indonesia) / 4

	if mean == 100 {
		return "Sempurna"
	} else if mean >= 90 {
		return "Sangat Baik"
	} else if mean >= 80 { 
		return "Baik"
	} else if mean >= 70 {
		return "Cukup"
	} else if mean >= 60 {
		return "Kurang"
	} else {
		return "Sangat kurang"
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 80, 100, 60))
}

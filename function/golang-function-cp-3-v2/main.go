package main

import (
	"fmt"
	"strings"
)

func FindShortestName(names string) string {
	tndbaca := []string{","," ", ";"}
	for _, tanda := range tndbaca {
		names = strings.ReplaceAll(names, tanda, ";")
	}

	nameList := strings.Split(names, ";") 
	shortestName := strings.TrimSpace(nameList[0])
	
	for _, name := range nameList {
		name := strings.TrimSpace(name)
		if len(name) < len(shortestName) || (len(name) == len(shortestName) && name < shortestName) {
			shortestName = name
		}
	}


	// for i := 0; i < len(cName); i++ {
	// 	fmt.Println(string(cName[i]))
	// }
	// var name string
	// for i := 0; i < len(names); i++ {
	// 	name = name + string(names[i])
	// 	if string(names[i]) == " " {
	// 		if len(name) == 2 {
	// 			return name
	// 		}
	// 		name = ""
	// 	}
	// }
	// fmt.Println(name)

	return shortestName // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	// fmt.Println(FindShortestName("Hanif;Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
}

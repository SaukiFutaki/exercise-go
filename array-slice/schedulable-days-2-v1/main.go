package main

import "fmt"

func SchedulableDays(villager [][]int) []int {
	// return nil // TODO: replace this
	var res []int
	if len(villager) == 0 {
		return []int{}
	}
	for i := 0; i < len(villager); i++ {
		if i == 0 {
			res = villager[i]
		} else {
			var t []int
			for j := 0; j < len(villager[i]); j++ {
				for k := 0; k < len(res); k++ {
					if villager[i][j] == res[k] {
						t = append(t, villager[i][j])
					}
				}
			}
			res = t
		}
	}
	return res
}


func main(){
	data := [][]int{{7, 12, 19, 22}, {12, 19, 21, 23}, {7, 12, 19}, {12, 19}}
	fmt.Println(SchedulableDays(data)) 
} 
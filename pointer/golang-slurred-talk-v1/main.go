package main

// import "fmt"

func SlurredTalk(words *string) {
	// TODO: answer here
	res := ""
	for _, w := range *words {
		// res += string(w)
		if w == 'S' || w == 'R' || w == 'Z' {
			res += "L"
		} else if w == 'r' || w == 'z' || w == 's' {
			res += "l"
		} else {
			res += string(w)
		}
	}
	*words = res
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Semangat pagi semuanya, semoga sehat selalu. Sehingga selalu bisa senyum"
	

	SlurredTalk(&words)
	
}

package main

import "fmt"

func EmailInfo(email string) string {
	var domain string
	var tld string

	index := -1
	for i := 0; i < len(email); i++ {
		
		fmt.Print(string(email[i]))
		if email[i] == '@' {
			// domain = email[i+1:]
			// break
			index = i
			break
		}


	}

	dotIndex := -1
	for i := index+1; i < len(email); i++ {
		fmt.Print(string(email[i]))
		if email[i] == '.' {
			dotIndex = i
			break
		}
	}

	domain = email[index+1:]
	tld = email[dotIndex+1:]
	return "domain: " + domain + " tld: " + tld + " index: " + string(index)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}

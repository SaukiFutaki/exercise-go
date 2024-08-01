package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	var totalPrice = (VIP * 30) + (regular * 20) + (student * 10)
	var totalTicket = VIP + regular + student


	if totalPrice > 100 {
		if day % 2 == 1 {
			if totalTicket < 5 {
				totalPrice = totalPrice - (15*totalPrice)/100
			} else if totalTicket > 5 {
				totalPrice = totalPrice - (25*totalPrice) / 100
			}
		} else if day % 2 == 0 {
			if totalTicket < 5 {
				totalPrice = totalPrice - (10 * totalPrice) / 100
			} else if totalTicket >5 {
				totalPrice = totalPrice - (20 *totalPrice) / 100
			}
		}
		
	}
	return float32(totalPrice)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(3, 3, 3, 20))
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here

func DeliveryOrder(data []string, day string) map[string]float32 {
	result := make(map[string]float32)
	for _, d := range data {
		// fmt.Println(d)
		tokens := strings.Split(d, ":")
		// fmt.Println(tokens)
		fName := tokens[0]
		lName := tokens[1]
		price, _ := strconv.ParseFloat(tokens[2], 32)
		city := tokens[3]
		// fmt.Println("First Name: ", fName)
		// fmt.Println("Last Name: ", lName)
		// fmt.Println("Price: ", price)
		// fmt.Println("City: ", city)

		if city == "BDG" {
			if !(day == "rabu" || day == "kamis" || day == "sabtu") {
				continue
			}
		}

		if city == "JKT" {
			if day == "minggu" {
				continue
			}
		}

		if city == "BKS" {
			if !(day == "selasa" || day == "kamis" || day == "jumat") {
				continue
			} 
		}
		if city == "DPK" {
			if !(day == "senin" || day == "selasa") {
				continue
			}
		}

		if day == "senin" || day == "rabu" || day == "jumat" {
			price += price * 0.1
		} else if day == "selasa" || day == "kamis" || day == "sabtu" {
			price += price * 0.05
		}

		result[fName+"-"+lName] += float32(price)

	

	}
	// return nil // TODO: replace this
	fmt.Println(result)
	return result
}

func main() {
	data := []string{
		"Budi:Gunawan:10000:JKT",
		"Andi:Sukirman:20000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Budi:Gunawan:50000:DPK",
	}

	day := "sabtu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}

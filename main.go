package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = conferenceTickets

	fmt.Printf("Welcome to %v booking applications\n", conferenceName)
	fmt.Println("Get your tickets to attend here")
	fmt.Printf("We have a total of %v tickets and %v are still remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings []string

	for remainingTickets > 0 && len(bookings) < 50 {

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enetr your first name")
		fmt.Scan(&firstName)

		fmt.Println("Enetr your last name")
		fmt.Scan(&lastName)

		fmt.Println("Enetr your email")
		fmt.Scan(&email)

		fmt.Println("Enetr number of tickets name")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("The whole array: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("Array type: %T\n", bookings)
			fmt.Printf("Array length: %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booling %v tickets. You will receive a confirmation email %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("Remaining tickets are %v\n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These first names of bookings are: %v\n", firstNames)

			noTicketsRemainig := remainingTickets == 0
			if noTicketsRemainig {
				fmt.Println("The tickets are sold out, Come back next year")
				break
			}

		} else {
			fmt.Printf("We ony have %v tickets remining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}
	}
}

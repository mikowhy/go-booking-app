package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have total of %v we still have %v available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email: ")
		fmt.Scan(&email)

		fmt.Print("How many ticket do you need?: ")
		fmt.Scan(&userTickets)

		if userTickets < remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank your %v %v for booking %v tickets. You will recive confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets have left", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("Theses are first names of booking %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out.")
				break
			}
		} else {
			fmt.Printf("We have only %v tickets. Try again\n", remainingTickets)
		}
	}
}

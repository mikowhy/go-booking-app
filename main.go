package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings []string

func main() {

	sayHello()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isNameValid, isEmailValid, isuserTicketsValid := validateUserInput(firstName, lastName, email, userTickets)

		if isNameValid && isEmailValid && isuserTicketsValid {
			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("Theses are first names of booking %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out.")
				break
			}
		} else {
			if !isNameValid {
				fmt.Println("First or last name is invalid")
			}

			if !isEmailValid {
				fmt.Println("Email has to have '@'")
			}

			if !isuserTicketsValid {
				fmt.Println("Ticket number is invalid")
			}
		}
	}
}

func sayHello() {
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have total of %v we still have %v available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isNameValid := len(firstName) >= 2 && len(lastName) >= 2
	isEmailValid := strings.Contains(email, "@")
	isUserTicketsValid := userTickets > 0 && userTickets <= remainingTickets

	return isNameValid, isEmailValid, isUserTicketsValid
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank your %v %v for booking %v tickets. You will recive confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets have left\n", remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

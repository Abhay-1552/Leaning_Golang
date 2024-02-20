package main

import (
	"fmt"
	"strings"
)

func main() {

	// Use of variables
	var conferenceName = "Go Conference"
	const conferenceTicket = 50
	var remainingTickets uint = 50 // uint is an unsigned integer
	var bookings []string          // Array of strings of undefined length

	fmt.Println("Welcome to the", conferenceName, "booking system.")
	fmt.Println("We have", remainingTickets, "tickets remaining for the", conferenceName, "out of", conferenceTicket, "tickets.")

	for {
		// Declare variables
		var firstName string
		var lastName string
		var email string
		var userTicket uint

		// Ask user for their name and the number of tickets they want to book
		fmt.Print("\nEnter your First Name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your Last Name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your Email: ")
		fmt.Scan(&email)

		fmt.Print("Enter the number of tickets you want to book: ")
		fmt.Scan(&userTicket)

		// If the user tries to book more tickets than are available, ask them to try again
		if userTicket <= remainingTickets {
			fmt.Println("\nThank you", firstName, lastName, "for booking", userTicket, "tickets for the", conferenceName)

			// reduce the number of tickets
			remainingTickets = remainingTickets - userTicket

			fmt.Println("We have", remainingTickets, "tickets remaining for the", conferenceName, "out of", conferenceTicket, "tickets.")

			// Add the user to the bookings array
			var fullName = firstName + " " + lastName
			bookings = append(bookings, fullName)

			firstNames := []string{}
			// _ is a blank identifier used to ignore the index
			for _, name := range bookings {
				var first = strings.Fields(name)
				firstNames = append(firstNames, first[0])
			}

			fmt.Println("The following people have booked tickets for the", conferenceName, ":", firstNames)

		} else {
			fmt.Println("Sorry, there are only", remainingTickets, "tickets remaining and you have tried to book", userTicket, "tickets.")
			continue
		}

		// If there are no tickets remaining, stop the loop
		if remainingTickets == 0 {
			fmt.Println("Sorry, there are no tickets remaining for the", conferenceName)
			break
		}
	}
}

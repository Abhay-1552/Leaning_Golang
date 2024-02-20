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

	// Use to repeat a string multiple times
	fmt.Println(strings.Repeat("-", 75))

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

		// Validate the user's input
		var validName = len(firstName) > 2 && len(lastName) > 2

		// Use to check if a string contains a substring
		var validEmail = strings.Contains(email, "@") && strings.Contains(email, ".")

		var validTicket = userTicket > 0 && userTicket <= remainingTickets

		// If the user data is valid, book the tickets
		if validName && validEmail && validTicket {
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

				// Use to split a string into substrings
				var first = strings.Fields(name)

				firstNames = append(firstNames, first[0])
			}

			fmt.Println("\nThe following people have booked tickets for the", conferenceName, ":", firstNames)

		} else {
			if !validName {
				fmt.Println("\nSorry, your name is invalid. Please try again.")
			}
			if !validEmail {
				fmt.Println("\nSorry, your email is invalid. Please try again.")
			}
			if !validTicket {
				fmt.Println("\nSorry, the number of tickets you want to book is invalid. Please try again.")
			}
			continue
		}

		// If there are no tickets remaining, stop the loop
		if remainingTickets == 0 {
			fmt.Println("\nSorry, there are no tickets remaining for the", conferenceName)
			break
		}
	}
}

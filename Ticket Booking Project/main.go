package main

import "fmt"

func main() {

	// Use of variables
	var conferenceName = "Go Conference"
	const conferenceTicket = 50
	var remainingTickets uint = 50 // uint is an unsigned integer
	var bookings []string          // Array of strings of undefined length

	fmt.Println("Welcome to the", conferenceName, "booking system.")
	fmt.Println("We have", remainingTickets, "tickets remaining for the", conferenceName, "out of", conferenceTicket, "tickets.")

	for i := 0; i < 2; i++ {
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

		fmt.Println("\nThank you", firstName, lastName, "for booking", userTicket, "tickets for the", conferenceName)

		// reduce the number of tickets
		remainingTickets = remainingTickets - userTicket

		fmt.Println("We have", remainingTickets, "tickets remaining for the", conferenceName, "out of", conferenceTicket, "tickets.")

		// Add the user to the bookings array
		var fullName = firstName + " " + lastName
		bookings = append(bookings, fullName)

		firstNames := []string{}
		firstNames = append(firstNames, firstName)

		fmt.Println("The following people have booked tickets for the", conferenceName, ":", firstNames)
	}

	// fmt.Println("\nThe following people have booked tickets for the", conferenceName, "conference:", bookings)
}

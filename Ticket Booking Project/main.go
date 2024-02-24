package main

import (
	"fmt"
	"strings"
)

const conferenceTicket = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50 // uint is an unsigned integer
var bookings []string          // Array of strings of undefined length

func main() {
	// function to greet the user
	greetUser()

	for {
		// function to get user input
		var firstName, lastName, email, userTicket = getUserInput()

		// function to validate user input
		var isValidName, isValidEmail, isValidTicket = validateUser(firstName, lastName, email, userTicket)

		// If the user data is valid, book the tickets
		if isValidName && isValidEmail && isValidTicket {
			// function to book tickets
			bookTickets(firstName, lastName, userTicket)

			// function to print the first name of the people who have booked tickets
			var firstNames = printFirstName()
			fmt.Println("\nThe following people have booked tickets for the", conferenceName, ":", firstNames)

			// If there are no tickets remaining, stop the loop
			if remainingTickets == 0 {
				fmt.Println("\nSorry, there are no tickets remaining for the", conferenceName)
				break
			}

		} else {
			if !isValidName {
				fmt.Println("\nSorry, your name is invalid. Please try again.")
			}
			if !isValidEmail {
				fmt.Println("\nSorry, your email is invalid. Please try again.")
			}
			if !isValidTicket {
				fmt.Println("\nSorry, the number of tickets you want to book is invalid. Please try again.")
			}
			continue
		}
	}
}

// Greet the user
func greetUser() {
	fmt.Println("Welcome to the", conferenceName, "booking system.")
	fmt.Println("We have", remainingTickets, "tickets remaining for the", conferenceName, "out of", conferenceTicket, "tickets.")

	// Use to repeat a string multiple times
	fmt.Println(strings.Repeat("-", 75))
}

// Get the user's input
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Print("\nEnter your First Name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your Last Name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your Email: ")
	fmt.Scan(&email)

	fmt.Print("Enter the number of tickets you want to book: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

// Validate the user's input
func validateUser(firstName string, lastName string, email string, userTicket uint) (bool, bool, bool) {
	var isValidName = len(firstName) > 2 && len(lastName) > 2

	// Use to check if a string contains a substring
	var isValidEmail = strings.Contains(email, "@") && strings.Contains(email, ".")

	var isValidTicket = userTicket > 0 && userTicket <= remainingTickets

	return isValidName, isValidEmail, isValidTicket
}

// function for booking tickets
func bookTickets(firstName string, lastName string, userTicket uint) ([]string, uint) {
	fmt.Println("\nThank you", firstName, lastName, "for booking", userTicket, "tickets for the", conferenceName)

	// reduce the number of tickets
	remainingTickets = remainingTickets - userTicket

	fmt.Println("We have", remainingTickets, "tickets remaining for the", conferenceName, "out of", conferenceTicket, "tickets.")

	// Add the user to the bookings array
	var fullName = firstName + " " + lastName
	bookings = append(bookings, fullName)

	return bookings, remainingTickets
}

// Print the first name of the people who have booked tickets
func printFirstName() []string {
	firstNames := []string{}

	// _ is a blank identifier used to ignore the index
	for _, name := range bookings {

		// Use to split a string into substrings
		var first = strings.Fields(name)

		firstNames = append(firstNames, first[0])
	}
	return firstNames
}

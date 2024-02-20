package main

import "fmt"

func main() {

	// Use of variables
	var conferenceName = "Go Conference"
	const conferenceTicket = 50
	var remainingTickets = 50

	fmt.Println("Welcome to the", conferenceName, "booking system.")
	fmt.Println("We have", remainingTickets, "tickets remaining for the", conferenceName, "conference out of", conferenceTicket, "tickets.")

	// User details
	var userName string
	var userTicket int

	// Declared values for variables
	userName = "John"
	userTicket = 5

	fmt.Println("Hello", userName, "you have booked", userTicket, "tickets for the", conferenceName, "conference.")
}

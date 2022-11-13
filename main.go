package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	var conferenceTickets int = 100
	var remainingTickets int = conferenceTickets
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend the %v\n", conferenceName)


	var username string
	var userTickets int

	username = "kalisa"
	userTickets = 2	
	fmt.Printf("User %v booked %v tickets.\n", username, userTickets)

}

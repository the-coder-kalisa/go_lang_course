package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	var conferenceTickets uint = 100
	var remainingTickets uint = conferenceTickets
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend the %v\n", conferenceName)

	var firstname string
	var lastname string
	var email string
	var userTickets uint
	fmt.Print("Enter your firstname: ")
	fmt.Scan(&firstname)
	fmt.Print("Enter your lastname: ")
	fmt.Scan(&lastname)
	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	remainingTickets -= userTickets
	fmt.Printf("Thank you %v %v for booking %v tickets. You'll recieve a confirmation email at %v\n", firstname, lastname, userTickets, email)
}

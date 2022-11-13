package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	bookings := []string{}
	var conferenceTickets uint = 100
	var remainingTickets uint = conferenceTickets
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend the %v\n", conferenceName)

	for {

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
		if userTickets <= remainingTickets {

			remainingTickets -= userTickets

			bookings = append(bookings, firstname+" "+lastname)

			fmt.Printf("Thank you %v %v for booking %v tickets. You'll recieve a confirmation email at %v\n", firstname, lastname, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
			firstnames := []string{}
			for _, bookings := range bookings {
				var names = strings.Fields(bookings)
				var firstName = names[0]

				firstnames = append(firstnames, firstName)

			}
			fmt.Printf("These are firstnames of  the bookings in the application: %v\n", firstnames)
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. come back next year.")
				break
			}
		} else {
			fmt.Printf("We only have %v remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}
	}
}

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
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

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
		isValidEmail, isValidName, isValidTicketNumber := validateInput(firstname, lastname, email, userTickets, remainingTickets)
		if isValidEmail && isValidName && isValidTicketNumber {

			remainingTickets -= userTickets

			bookings = append(bookings, firstname+" "+lastname)

			fmt.Printf("Thank you %v %v for booking %v tickets. You'll recieve a confirmation email at %v\n", firstname, lastname, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
			firstnames := printFirstnames(bookings)
			fmt.Printf("These are firstnames of  the bookings in the application: %v\n", firstnames)
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("firstname or lastname your entered is too short")
			} else if !isValidEmail {
				fmt.Println("email addres you entered is doesn't contain @ in it")
			} else if !isValidTicketNumber {
				fmt.Println("number of ticket you entered is invalid")
			}
			continue
		}
	}
}

func greetUsers(conferenceName string, conferenceTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend the %v\n", conferenceName)
}

func printFirstnames(bookings []string) []string {
	firstnames := []string{}
	for _, bookings := range bookings {
		var names = strings.Fields(bookings)
		var firstName = names[0]

		firstnames = append(firstnames, firstName)

	}
	return firstnames
}

func validateInput(firstname string, lastname string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(lastname) >= 2 && len(firstname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidEmail, isValidName, isValidTicketNumber
}

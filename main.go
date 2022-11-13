package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Go Conference"
var bookings = make([]userData, 0)
var conferenceTickets uint = 100
var remainingTickets uint = conferenceTickets

type userData struct {
	firstname       string
	lastname        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()
	for {

		firstname, lastname, email, userTickets := getUserInputs()

		isValidEmail, isValidName, isValidTicketNumber := helper.ValidateInput(firstname, lastname, email, userTickets, remainingTickets)
		if isValidEmail && isValidName && isValidTicketNumber {
			bookTicket(userTickets, firstname, lastname, email)
			wg.Add(1)
			go sendTicket(userTickets, firstname, lastname, email)
			firstnames := printFirstnames()
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
		wg.Wait()
	}

}
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend the %v\n", conferenceName)
}

func printFirstnames() []string {
	firstnames := []string{}
	for _, bookings := range bookings {
		var firstName = bookings.firstname

		firstnames = append(firstnames, firstName)

	}
	return firstnames
}

func getUserInputs() (string, string, string, uint) {
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

	return firstname, lastname, email, userTickets
}

func bookTicket(userTickets uint, firstname string, lastname string, email string) {
	remainingTickets -= userTickets

	// create map to inser user

	var userData = userData{
		firstname:       firstname,
		lastname:        lastname,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You'll recieve a confirmation email at %v\n", firstname, lastname, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstname string, lastname string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstname, lastname)
	fmt.Println("##################")
	fmt.Printf("Sending ticket:'n %v to email address %v\n", ticket, email)
	fmt.Println("##################")
	wg.Done()
}

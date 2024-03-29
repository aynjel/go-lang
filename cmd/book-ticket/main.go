package main

import (
	"fmt"
	"sync"
	"test-app/pkg/helper"
	"time"
)

const conferenceTickets int = 100

var conferenceName = "Go conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUser()

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTicketNumber := getUserInput()

		isValidName, isValidEmail, isValidUserTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTicketNumber, remainingTickets)

		if isValidName && isValidEmail && isValidUserTicketNumber {
			bookTicket(userTicketNumber, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTicketNumber, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("Attendees: %s\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("Sorry, %s is sold out!\n", conferenceName)
				// break
			}
		} else {
			if !isValidName {
				fmt.Println("Please enter a valid name")
			}
			if !isValidEmail {
				fmt.Println("Please enter a valid email")
			}
			if !isValidUserTicketNumber {
				fmt.Println("Please enter a valid number of tickets")
			}
		}
		wg.Wait()
	}
}

func greetUser() {
	fmt.Printf("Welcome to %s, we have %d tickets remaining!\n", conferenceName, remainingTickets)
	fmt.Printf("We have %d tickets remaining!\n", remainingTickets)
	fmt.Println("Get your tickets now before they sell out!")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicketNumber uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("How many tickets would you like to purchase?")
	fmt.Scan(&userTicketNumber)

	return firstName, lastName, email, userTicketNumber
}

func bookTicket(userTicketNumber uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTicketNumber
	userData := UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   userTicketNumber,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thanks for purchasing tickets, %s! You now have %d tickets.\n", firstName, userTicketNumber)
	fmt.Printf("There are now %d tickets remaining.\n", remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(2 * time.Second)
	ticket := fmt.Sprintf("Hi %s, thanks for purchasing %d tickets to %s!", firstName, userTickets, conferenceName)
	fmt.Printf("Sending ticket to %s...\n", email)
	fmt.Println(ticket)
	wg.Done()
}

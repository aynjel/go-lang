package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTicketNumber uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 3 && len(lastName) >= 3
	isValidEmail := strings.Contains(email, "@")
	isValidUserTicketNumber := userTicketNumber > 0 && userTicketNumber <= remainingTickets

	return isValidName, isValidEmail, isValidUserTicketNumber
}

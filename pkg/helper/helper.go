package helper

import (
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, userTicketNumber uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 3 && len(lastName) >= 3
	isValidEmail := strings.Contains(email, "@")
	isValidUserTicketNumber := userTicketNumber > 0 && userTicketNumber <= remainingTickets

	return isValidName, isValidEmail, isValidUserTicketNumber
}

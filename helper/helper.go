package helper

import (
	"fmt"
	"strings"
	"time"
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func ValidateUserInput(firstName, lastName, email string, userTickets, remainingTickets uint) (bool, bool, bool) {
	isValidNames := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidNames, isValidEmail, isValidTicketNumber
}

func GetFirstNames(bookings []UserData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func GetUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// Ask user for their name
	fmt.Println("Enter your first name.")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name.")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address.")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets.")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func BookTicket(remainingTickets, userTickets uint, bookings []UserData, firstName, lastName, email, conferenceName string) ([]UserData, uint) {
	remainingTickets = remainingTickets - userTickets

	// Create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return bookings, remainingTickets
}

func SendTicket(userTickets uint, firstName, lastName, email string) {
	// Simulate delay to generate the tickets
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##############")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("##############")
}

package main

import (
	"fmt"

	"github.com/ciscodea/booking-app/helper"
)

func greetUsers(conferenceName string, conferenceTickets, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets
	var bookings = make([]helper.UserData, 0)

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		firstName, lastName, email, userTickets := helper.GetUserInput()
		isValidNames, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidNames && isValidEmail && isValidTicketNumber {
			bookings, remainingTickets = helper.BookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

			// Start a goroutine to generate Tickets concurrently
			go helper.SendTicket(userTickets, firstName, lastName, email)

			fmt.Printf("The first names of bookings are: %v\n", helper.GetFirstNames(bookings))

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidNames {
				fmt.Println("Your name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Your email address you entered doesn't contain @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered are invalid.")
			}
		}
	}
}

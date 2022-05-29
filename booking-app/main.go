package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName string = "Go Conference"

var remainingTickets int = conferenceTickets

var bookings []string

func main() {

	greetUsers()

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// fmt.Println(conferenceName)

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			// call function to print first names
			// printFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", printFirstNames())

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conferece is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("The Email address you entered does not contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
			continue
		}

	}
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Println("We have total of ", conferenceTickets, "tickets and", remainingTickets, "are still available")

	fmt.Println("Book your tickets in here for:")
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])

	}
	fmt.Print(firstNames)
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets you want?  ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return
}

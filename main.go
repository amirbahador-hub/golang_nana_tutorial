package main // a package is a collection of GO files

import (
	"fmt"
	"time"
)

var conferenceName = "GO Conference"

const conferenceTickets = 50 // Not Changeable
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // now this is an slice

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func main() {

	greetUser()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidEmail, isValidName, isValidTicketsNumber := ValidatedUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketsNumber {

			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("The first names is %v \n", firstNames)

			if remainingTickets == 0 {
				// end the program
				fmt.Println("Our conference is booked out. see you next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first Name or last name is short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't have @ sign")
			}
			if !isValidTicketsNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			fmt.Printf("your input is invalid, TRY AGAIN  \n")

		}
	}

}

func greetUser() {
	fmt.Printf("Welcome to %v booking application", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking) // split separate by space
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for user
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v \n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceTickets)
}


func sendTicket(userTickets uint, firstName string, lastName string, email string){

	time.Sleep(10 * time.Second) // simulation of log task
	// for example generate pdf and send to email by real

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###########################")
	fmt.Printf("Sending ticket: \n %v \nto email address %v\n", ticket, email)
	fmt.Println("###########################")
}

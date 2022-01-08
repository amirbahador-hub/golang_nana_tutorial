package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "GO Conference"
	const conferenceTickets = 50 // Not Changeable
	var remainingTickets uint = 50
	bookings := []string{} // now this is an slice

	greetUser(conferenceName, conferenceTickets , remainingTickets)

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidEmail, isValidName, isValidTicketsNumber := validatedUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketsNumber {
 
			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)
			firstNames := getFirstNames(bookings)
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

func greetUser(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames(bookings[]string) []string {

	firstNames := []string{}
	for _ , booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames

}

func validatedUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
		isValidName :=  len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketsNumber := userTickets > 0 && userTickets <= remainingTickets

		return isValidName, isValidEmail, isValidTicketsNumber
}

func getUserInput () (string, string, string, uint){
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

func bookTicket (remainingTickets uint, userTickets uint, bookings []string ,firstName string, lastName string, email string ,conferenceTickets string){
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceTickets)
	
}
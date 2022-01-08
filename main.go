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

	fmt.Printf("conferenceTickets is %T, remainingTickets %T, conferenceName is %T \n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for remainingTickets>0 && len(bookings) < 50{

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

		if userTickets < remainingTickets {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceTickets)
			
			firstNames := []string{}
			for _ , booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of  bookings are : %v \n", firstNames)

			if remainingTickets == 0 {
				// end the program
				fmt.Println("Our conference is booked out. see you next year")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets so your can't book %v tickets \n", remainingTickets, userTickets)
			
		}
		
	}
}






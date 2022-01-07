package main
import "fmt"

func main() {
	var conferenceName = "GO Conference"
	const conferenceTickets = 50        // Not Changeable
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")


}


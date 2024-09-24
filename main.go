package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

// Conference-related information
var conferenceName = "Go Conference"
const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)
var bookingMutex sync.Mutex  // Mutex to prevent data race on bookings

// Struct to store user booking data
type UserData struct {
	FirstName       string
	LastName        string
	Email           string
	NumberOfTickets uint
}

func main() {
	// Initialize logger to capture important events
	log.SetPrefix("Booking System: ")
	log.SetFlags(0)

	// Display a greeting to the user
	greetUser()

	// Loop for continuous input until tickets are sold out or user exits
	for {
		// Get user input for booking
		firstName, lastName, email, userTickets := getUserInput()

		// Validate user input
		isValid, err := validateUserInput(firstName, lastName, email, userTickets)
		if !isValid {
			log.Println(err)
			continue  // Skip to the next iteration if validation fails
		}

		// Book tickets and send confirmation asynchronously
		bookTickets(userTickets, firstName, lastName, email)
		go sendTickets(userTickets, firstName, lastName, email)

		// Display current bookings
		printFirstNames()

		// End loop if tickets are sold out
		if remainingTickets == 0 {
			fmt.Println("All tickets have been sold out. Come back next year.")
			break
		}
	}
}

// Greets the user with conference information
func greetUser() {
	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
}

// Displays the first names of people who have booked tickets
func printFirstNames() {
	bookingMutex.Lock()  // Prevent race conditions on bookings slice
	defer bookingMutex.Unlock()

	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.FirstName)
	}
	fmt.Printf("First names of bookings: %v\n", firstNames)
}

// Get user input for the booking process
func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	// Collect user details
	fmt.Print("Please enter your first name: ")
	fmt.Scanln(&firstName)

	fmt.Print("Please enter your last name: ")
	fmt.Scanln(&lastName)

	fmt.Print("Please enter your email address: ")
	fmt.Scanln(&email)

	fmt.Print("How many tickets would you like to purchase? ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

// Book tickets for the user and update the remaining ticket count
func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	bookingMutex.Lock()
	defer bookingMutex.Unlock()

	remainingTickets -= userTickets  // Deduct the booked tickets from available tickets

	// Create a booking instance and append it to bookings slice
	var userData = UserData{
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		NumberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)

	// Provide feedback to the user
	fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email will be sent to %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining.\n", remainingTickets)
}

// Sends the ticket confirmation via email simulation (concurrent function)
func sendTickets(userTickets uint, firstName, lastName, email string) {
	time.Sleep(10 * time.Second)  // Simulate delay in sending tickets
	confirmation := fmt.Sprintf("%v ticket(s) sent to %v %v at %v", userTickets, firstName, lastName, email)
	fmt.Println("#####################")
	fmt.Println(confirmation)
	fmt.Println("#####################")
}

// Validate user input for correctness
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, error) {
	// Ensure names have at least 2 characters
	if len(firstName) < 2 || len(lastName) < 2 {
		return false, fmt.Errorf("invalid name: first and last names must have at least 2 characters")
	}
	// Check if the email contains an "@" symbol
	if !strings.Contains(email, "@") {
		return false, fmt.Errorf("invalid email address")
	}
	// Ensure the user is booking at least one ticket and not exceeding the available tickets
	if userTickets <= 0 || userTickets > remainingTickets {
		return false, fmt.Errorf("invalid ticket number: must be between 1 and %v", remainingTickets)
	}

	// Input is valid
	return true, nil
}

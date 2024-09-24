
```markdown
# GoLang Conference Tickets Booking Application

This is a simple ticket booking system for a fictional "Go Conference" built using GoLang. The application allows users to book tickets, validate input, and handle bookings concurrently, simulating email confirmations with delays.

## Features

- **User Input**: Users can input their name, email, and the number of tickets they want to purchase.
- **Validation**: The system validates user input such as name length, valid email, and ticket availability.
- **Concurrency**: Ticket confirmations are sent asynchronously, simulating real-world email delays.
- **Thread Safety**: The booking system is thread-safe, using a mutex to avoid race conditions on shared resources.

## How It Works

1. Users are greeted with available tickets and prompted to enter their details.
2. User input is validated for correctness.
3. Upon successful validation, tickets are booked, and a confirmation message is displayed.
4. The application sends a simulated email confirmation with a 10-second delay.
5. The application runs in a loop until all tickets are sold out.

## Project Structure

```
GoLangConferenceTicketsBookingApplication/
│
├── main.go         # Main application logic
├── README.md       # Project documentation
└── go.mod          # Go module definition
```

### `main.go`

This file contains the main logic for the ticket booking system:
- **`greetUser()`**: Greets the user with the conference details.
- **`getUserInput()`**: Collects user details like first name, last name, email, and number of tickets.
- **`validateUserInput()`**: Ensures valid user input such as valid email and ticket availability.
- **`bookTickets()`**: Books tickets for the user and updates the remaining ticket count.
- **`sendTickets()`**: Simulates sending ticket confirmation via email with a delay.
- **Concurrency Handling**: Mutex is used to prevent race conditions on shared resources (i.e., bookings).

## Requirements

- **GoLang**: Make sure you have Go installed on your machine. [Download Go](https://golang.org/dl/)

## Running the Application

1. Clone the repository:

```bash
git clone https://github.com/Abdul-Firoz-Khan/GoLangConferenceTicketsBookingApplication.git
```

2. Navigate to the project directory:

```bash
cd GoLangConferenceTicketsBookingApplication
```

3. Run the application:

```bash
go run main.go
```

## Example

```
Welcome to the Go Conference booking application
We have a total of 50 tickets and 50 are still available.

Please enter your first name: John
Please enter your last name: Doe
Please enter your email address: john.doe@example.com
How many tickets would you like to purchase? 2
Thank you John Doe for booking 2 tickets. A confirmation email will be sent to john.doe@example.com.
48 tickets remaining.

First names of bookings: [John]
#####################
2 ticket(s) sent to John Doe at john.doe@example.com
#####################
```

## License

This project is licensed under the MIT License.
```

This `README.md` file provides a comprehensive overview of the project, setup instructions, and an example of how the application works. Let me know if you'd like any changes!

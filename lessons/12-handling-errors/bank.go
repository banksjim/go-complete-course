package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Set global constant names for input/output files
const accountBalanceFile = "AccountBalance.txt"

func getBalanceFromFile() float64 {

	// Read contents of the account balance file
	data, _ := os.ReadFile(accountBalanceFile) // Use the _ when you want to ignore the returned value 

    // Parse the values from the accountBalanceFile
	balanceText := string(data)
	parts := strings.Split(balanceText, ",")
	// dateTimeString := parts[0]
	balanceString := strings.TrimSpace(parts[1]) // Remove any whitespace or newlines

	// Convert balance string to float64
	balance, _ := strconv.ParseFloat(balanceString, 64)

	// Only return the balance amount for now. 
	// We could also return and use the dateTimeString later.
	return balance
}

func writeBalanceToFile(balance float64) {

    // Get the current date/time
	currentDateTime := time.Now()
	
	// Form the string to write to the AccountBalance file
	balanceText := fmt.Sprintf("%s,%.2f\n", currentDateTime, balance)

    // Output to the AccountBalance.txt file
	os.WriteFile(accountBalanceFile, []byte(balanceText),0644)

}

func main() {
 
    // Declare function variables
	userAmount     := 0.0
	userSelection  := 0

	// Retrieve the initial account balance
	accountBalance := getBalanceFromFile()

	// Present welcome message
	fmt.Println("\nWelcome to Go Bank!")

    // Loop until user selects the Exit option
	for {

		// Present app options
		fmt.Println("\nWhat do you want to do?")
		fmt.Println("\n1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		// Select user options
		fmt.Print("\nSelection: ")
		fmt.Scan(&userSelection)

		switch userSelection {

		case 1:

			// Print account balance
			fmt.Printf("Account balance: %.2f\n", accountBalance)

		case 2:

			// Enter deposit amount
			fmt.Print("Deposit amount: ")
			fmt.Scan(&userAmount)

            // Validation guard checks:
			// Check: Deposit amount greater than 0
			if userAmount <= 0 {

                fmt.Println("Invalid amount. Must be greater than zero.")
				continue

			}

            // Add deposit amount to balance
			accountBalance += userAmount

			// Output updated account balance to the AccountBalance file
			writeBalanceToFile(accountBalance)

		case 3:

			// Enter withdrawal amount
			fmt.Print("Withdrawal amount: ")
			fmt.Scan(&userAmount)

            // Validation guard checks:
			// Check: Withdrawal amount must be greater than zero
			if userAmount <= 0 {

                fmt.Println("Invalid amount. Must be less than or equal to account balance.")
				continue

			}

            // Check: Withdrawal amount must be less than or equal balance
			if userAmount > accountBalance {

                fmt.Println("Invalid amount. Must be less than or equal to account balance.")
				continue

			}

			// Subtract withdrawal amount from balance
			accountBalance -= userAmount

			// Output updated account balance to the AccountBalance file
			writeBalanceToFile(accountBalance)

		case 4:

	        // Exit the main function
			fmt.Println("Thanks for choosing our bank!")
			return

		default:

			// Error: User did not select a valid menu option
			fmt.Println("**Input error**")

		} // End switch

	} // End menu control loop

} // End main()
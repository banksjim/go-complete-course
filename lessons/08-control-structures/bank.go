package main

import "fmt"

func main() {
 
    // Declare function variables
	accountBalance := 1000.00
	userAmount     := 0.0
	userSelection  := 0

	// Present welcome message
	fmt.Println("\nWelcome to Go Bank!")
	fmt.Println("What do you want to do?")

	// Present app options
	fmt.Println("\n1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

    // Loop until user selects the Exit option
	for {

		// Select user options
		fmt.Print("\nSelection: ")
		fmt.Scan(&userSelection)

		// Handle user options
		if userSelection >= 1 && userSelection <= 4 {
			
			// Check balance
			if userSelection == 1 {
				
				// Print account balance
				fmt.Printf("Account balance: %.2f\n", accountBalance)
			
			}

			// Deposit money
			if userSelection == 2 {

				// Enter deposit amount
				fmt.Print("Deposit amount: ")
				fmt.Scan(&userAmount)

				// Add deposit amount to balance
				accountBalance += userAmount

			}

			// Withdraw money
			if userSelection == 3 {

				// Enter withdrawal amount
				fmt.Print("Withdrawal amount: ")
				fmt.Scan(&userAmount)

				// Add deposit amount to balance
				accountBalance -= userAmount

			}

			// Exit - Terminate the app
			if userSelection == 4 {

				break;

			}

		} else {
			fmt.Println("**Input error**")
		}

    }

}
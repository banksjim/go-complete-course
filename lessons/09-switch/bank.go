package main

import "fmt"

func main() {
 
    // Declare function variables
	accountBalance := 1000.00
	userAmount     := 0.0
	userSelection  := 0

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
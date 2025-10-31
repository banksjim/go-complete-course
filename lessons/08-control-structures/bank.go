package main

import "fmt"

func main() {
 
    // Declare function variables
	var userSelection int = 0

	// Present welcome message
	fmt.Println("\nWelcome to Go Bank!")
	fmt.Println("What do you want to do?\n")

	// Present app options
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit\n")

    // Select user options
	fmt.Print("Selection: ")
	fmt.Scan(&userSelection)



}
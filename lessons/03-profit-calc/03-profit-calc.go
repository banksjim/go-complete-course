package main

import (
	"fmt"
	"math"
)

func main() {

    // Declare function variables
	expensesAmount := 0.0
	grossEarnings  := 0.0
	netEarnings    := 0.0
	revenueAmount  := 0.0
	taxRate        := 0.0

	// Request user inputs
	fmt.Print("Input revenue amount ($): ")
	fmt.Scan(&revenueAmount)

    fmt.Print("Input expenses amount ($): ")
	fmt.Scan(&expensesAmount)

	fmt.Print("Input tax rate (%): ")
	fmt.Scan(&taxRate)

	// Calculate gross earnings (earnings before taxes (EBT))
	grossEarnings = revenueAmount - expensesAmount

	// Calculate net earnings (earnings after taxes (profit))
	netEarnings = grossEarnings - (grossEarnings * (taxRate / 100))

	// Output results
	fmt.Printf("\nGross earnings:                    $ %.2f\n", (math.Ceil(grossEarnings * 100) / 100))
    fmt.Printf("Net earnings after taxes (profit): $ %.2f\n", (math.Ceil(netEarnings * 100) / 100))
}

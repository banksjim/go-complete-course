package main

import (
	"fmt"
	"math"
)

// Declare global constants
const InflationRate float64 = 2.67


func main() {
	
	// Declare main() variables
	var investmentAmount   float64 = 0
	var expectedReturnRate float64 = 0
	var formattedFV        string = ""
	var formattedRV        string = ""
	var yearsInvested      float64 = 0

    // Enter values to calculate
    fmt.Print("Input investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Input expected return rate percentage: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Input years of expected investment: ")
	fmt.Scan(&yearsInvested)

    // Calculate future and future real values
	futureValue, realValue := calculateFutureValues(investmentAmount, expectedReturnRate, yearsInvested)

    // Create formatted strings
	formattedFV = fmt.Sprintf("Future value: $%.2f\n", math.Round(futureValue * 100) / 100)
	formattedRV = fmt.Sprintf("Real value:   $%.2f\n", math.Round(realValue * 100) / 100)

	// Output future and real values using Sprintf() format specifiers
	fmt.Print(formattedFV)
	fmt.Print(formattedRV)

	// Output formatted strings using Printf()
	fmt.Printf("Future value: $%.2f\n", math.Round(futureValue * 100) / 100)
	fmt.Printf("Real value:   $%.2f\n", math.Round(realValue * 100) / 100)
}

func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, yearsInvested float64) (float64, float64) {

    fv  := investmentAmount * math.Pow(1 + (expectedReturnRate / 100), yearsInvested)
	rfv := fv / math.Pow(1 + (InflationRate/100), yearsInvested)

	return fv, rfv	       
}

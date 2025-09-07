package main

import (
	"fmt"
	"math"
)

func main() {
	
	// Declare main() constants
	const InflationRate    float64 = 2.67

	// Declare main() variables
	var investmentAmount   float64 = 0
	var expectedReturnRate float64 = 0
	var formattedFV        string = ""
	var formattedRV        string = ""
	var futureValue        float64 = 0
	var realValue          float64 = 0
	var yearsInvested      float64 = 0

    // Enter values to calculate
    fmt.Print("Input investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Input expected return rate percentage: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Input years of expected investment: ")
	fmt.Scan(&yearsInvested)

    // Calculate future value
	futureValue = investmentAmount * math.Pow(1 + (expectedReturnRate / 100), yearsInvested) 

	// Calculate the future real value (inflation adjusted)
	realValue = futureValue / math.Pow(1 + (InflationRate/100), yearsInvested)

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
 
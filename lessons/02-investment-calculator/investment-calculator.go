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
	var futureRealValue    float64 = 0
	var futureValue        float64 = 0
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
	futureRealValue = futureValue / math.Pow(1 + (InflationRate/100), yearsInvested)

	// Output future and real values
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
 
package main

import "fmt"
import "math"

func main() {
	
	// Declare main() variables
	var investmentAmount = 500000.00
	var expectedReturnRate = 8.5
	var yearsInvested = 10.0

    // Calculate future value
	var futureValue = float64(investmentAmount) * math.Pow((1 + (float64(expectedReturnRate) / 100)), float64(yearsInvested)) 

	// Output future value
	fmt.Println(futureValue)
}
 
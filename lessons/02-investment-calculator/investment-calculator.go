package main

import "fmt"
import "math"

func main() {
	
	// Declare main() variables
	var investmentAmount float64 = 500000.00
	var expectedReturnRate float64 = 8.5
	var futureValue float64 = 0.0
	var yearsInvested float64 = 10

    // Calculate future value
	futureValue = investmentAmount * math.Pow((1 + (expectedReturnRate / 100)), yearsInvested) 

	// Output future value
	fmt.Println(futureValue)
}
 
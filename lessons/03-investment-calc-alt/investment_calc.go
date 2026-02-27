package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate float64 = 2.5

	investmentAmount   := 1000.00
	expectedReturnRate := 5.50
	years              := 10.0

	futureValue                  := investmentAmount * math.Pow(1 + (expectedReturnRate / 100), years)
	inflationAdjustedFutureValue := futureValue / math.Pow(1 + (inflationRate / 100), years)

	fmt.Println("Future value:", futureValue)
	fmt.Println("Inflation adjusted future value:", inflationAdjustedFutureValue)
}
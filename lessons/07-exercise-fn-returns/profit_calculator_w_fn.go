package main

import "fmt"

func main() {

	revenue  := acceptInput("Revenue")
	expenses := acceptInput("Expenses")
	taxRate  := acceptInput("Tax Rate")

    ebt, profit, ratio := calcOutput(revenue, expenses, taxRate)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func acceptInput(valueName string) (float64) {
	var userInputValue float64 = 0
	
	fmt.Printf("%s: ", valueName)
	fmt.Scan(&userInputValue)

	return userInputValue
}

func calcOutput(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

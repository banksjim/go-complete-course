package main

import (
	"errors"
    "fmt"
	"os"
)

const profitOutputFile = "ProfitValues.txt"

func main() {

	// Input revenue
	revenue, err  := acceptInput("Revenue")

	if err != nil {
		
		fmt.Println(err)
		return // Exit main

	}

	// Input expenses
	expenses, err := acceptInput("Expenses")

    if err != nil {
		
		fmt.Println(err)
		return // Exit main

	}

    // Input tax rate
	taxRate, err := acceptInput("Tax Rate")

    if err != nil {
		
		fmt.Println(err)
		return // Exit main

	}

	// Call calcOutput to calculate values
    ebt, profit, ratio := calcOutput(revenue, expenses, taxRate)

    // Output calculated values to file
	writeProfitValuesFile(ebt, profit, ratio)

	// Output values to terminal
	fmt.Printf("%.2f\n", ebt)
	fmt.Printf("%.2f\n", profit)
	fmt.Printf("%.1f\n", ratio)
}

func acceptInput(valueName string) (float64, error) {
	
	var userInputValue float64 = 0
	
	fmt.Printf("%s: ", valueName)
	fmt.Scan(&userInputValue)

    if userInputValue <= 0 {
		
		return 0, errors.New("Value must be greater than zero")

	}

	return userInputValue, nil
}

func calcOutput(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

// Write values to file
func writeProfitValuesFile (ebtOutput float64, profitOutput float64, ratioOutput float64) {

	valuesOutput := fmt.Sprintf("%.2f,%.2f,%.2f\n", ebtOutput, profitOutput, ratioOutput)

	os.WriteFile(profitOutputFile, []byte(valuesOutput), 0644)
}


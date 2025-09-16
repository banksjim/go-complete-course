// Package main implements a profit calculator application that computes
// gross earnings, net earnings after taxes, and earnings ratio based on
// user-provided revenue, expenses, and tax rate inputs.
//
// This application demonstrates basic Go programming concepts including:
// - Variable declaration and initialization
// - User input collection using fmt.Scan
// - Mathematical calculations
// - Formatted output using fmt.Printf
// - Math package usage for ceiling operations
package main

import (
	"fmt"
	"math"
)

// main is the entry point of the profit calculator application.
// It collects user input for revenue, expenses, and tax rate, then
// calculates and displays:
// - Gross earnings (revenue minus expenses, before taxes)
// - Net earnings (profit after taxes)
// - Earnings ratio (relationship between gross and net earnings)
//
// The application uses interactive prompts to gather financial data
// and presents results with proper currency formatting.
func main() {

    // Declare and initialize financial calculation variables
    // All monetary values are stored as float64 for precision in calculations
	earningsRatio  := 0.0  // Ratio comparing gross to net earnings
	expensesAmount := 0.0  // Total business expenses in dollars
	grossEarnings  := 0.0  // Earnings before taxes (EBT) in dollars
	netEarnings    := 0.0  // Final profit after taxes in dollars
	revenueAmount  := 0.0  // Total revenue/income in dollars
	taxRate        := 0.0  // Tax rate as a percentage (e.g., 25.5 for 25.5%)

	// Interactive input collection phase
	// Collect revenue: total income before any deductions
	outputText("Input revenue amount ($): ")
	fmt.Scan(&revenueAmount)

    // Collect expenses: total business costs and expenditures
    outputText("Input expenses amount ($): ")
	fmt.Scan(&expensesAmount)

	// Collect tax rate: percentage of gross earnings paid as taxes
	outputText("Input tax rate (%): ")
	fmt.Scan(&taxRate)

	// Financial calculations phase
	// Step 1: Calculate gross earnings (Earnings Before Taxes)
	// Formula: Gross Earnings = Revenue - Expenses
	grossEarnings = revenueAmount - expensesAmount

	// Step 2: Calculate net earnings (profit after taxes)
	// Formula: Net Earnings = Gross Earnings - Tax Amount
	// Tax Amount = Gross Earnings × (Tax Rate ÷ 100)
	netEarnings = grossEarnings - (grossEarnings * (taxRate / 100))

	// Step 3: Calculate earnings ratio (efficiency metric)
	// This ratio shows the relationship between gross and net earnings
	// Higher ratios indicate more tax burden relative to net profit
	earningsRatio = grossEarnings / netEarnings

	// Results presentation phase
	// Display calculated financial metrics with proper formatting
	// Using math.Ceil() to round up to nearest cent for conservative estimates
	fmt.Printf("\nGross earnings:                    $ %.2f\n", (math.Ceil(grossEarnings * 100) / 100))
    fmt.Printf("Net earnings after taxes (profit): $ %.2f\n", (math.Ceil(netEarnings * 100) / 100))
	fmt.Printf("Earnings ratio:                      %.2f\n", (math.Ceil(earningsRatio * 100) / 100))
}

func outputText(textOut string) {
	fmt.Print(textOut)
}
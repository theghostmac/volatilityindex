package main

import (
	"fmt"
	"github.com/theghostmac/volatilityindex/utils"
	"github.com/theghostmac/volatilityindex/vix"
	"os"
)

func main() {
	// Log the current working directory.
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory: ", err)
		return
	}
	fmt.Println("Current working directory: ", cwd)

	// Read CSV file using utils package.
	filePath := "data/sp5000_index.csv"
	records, err := utils.ReadCSVFile(filePath)
	if err != nil {
		fmt.Println("Error reading CSV file: ", err)
		return
	}

	// Parse dates and values using utils package.
	dates, sp500Values, err := utils.ParseDateAndValue(records)
	if err != nil {
		fmt.Println("Error parsing dates and values: ", err)
		return
	}

	// Calculate returns and volatility
	returns := vix.CalculateReturns(sp500Values)
	volatility := vix.CalculateVolatility(returns)
	fmt.Printf("Volatility: %.2f%%\n", volatility*100)
}

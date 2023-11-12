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
	filePath := "data/sp500_index.csv"
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

	if len(dates) > 0 {
		fmt.Printf("Start Date: %s\n", dates[0].Format("2006-01-02"))
		fmt.Printf("End Date: %s\n", dates[len(dates)-1].Format("2006-01-02"))
	}

	// Calculate returns and volatility
	returns := vix.CalculateReturns(sp500Values)
	volatility := vix.CalculateVolatility(returns)
	fmt.Printf("Volatility: %.2f%%\n", volatility*100)

	//options := []data.Option{
	//	{StrikePrice: 100, Premium: 5},
	//	{StrikePrice: 110, Premium: 4},
	//}
	//riskFreeRate := 0.01   // 1%
	//timeToExpiration := 30 // 30 days
	//
	//vix := vix.CalculateVIX(options, riskFreeRate, float64(timeToExpiration))
	//fmt.Printf("VIX: %.2f\n", vix)
}

package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func CalculateReturns(prices []float64) []float64 {
	if len(prices) < 2 {
		fmt.Println("Error: Insufficient data for calculating returns.")
		return nil
	}

	returns := make([]float64, len(prices)-1)
	for i := 1; i < len(prices); i++ {
		returns[i-1] = (prices[i] - prices[i-1]) / prices[i-1]
	}
	return returns
}

func CalculateVolatility(returns []float64) float64 {
	if len(returns) == 0 {
		fmt.Println("Error: Insufficient data for calculating volatility.")
		return 0
	}

	var sum float64
	for _, ret := range returns {
		sum += ret
	}
	mean := sum / float64(len(returns))

	var sumSquaredDiff float64
	for _, ret := range returns {
		sumSquaredDiff += math.Pow(ret-mean, 2)
	}

	variance := sumSquaredDiff / float64(len(returns)-1)
	return math.Sqrt(variance)
}

func main() {
	// Log the current working directory.
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory: ", err)
		return
	}
	fmt.Println("Current working directory: ", cwd)

	// Open CSV file
	file, err := os.Open("data/sp500_index.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read CSV file
	reader := csv.NewReader(file)

	// Read the header row
	_, err = reader.Read()
	if err != nil {
		fmt.Println("Error reading header row: ", err)
		return
	}

	// Read remaining rows
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Extract dates and S&P 500 values
	var dates []time.Time
	var sp500Values []float64

	for _, record := range records {
		date, err := time.Parse("2006-01-02", record[0])
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		dates = append(dates, date)

		sp500, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			fmt.Println("Error parsing S&P 500 value: ", err)
			return
		}
		sp500Values = append(sp500Values, sp500)
	}

	// Calculate returns and volatility
	returns := CalculateReturns(sp500Values)
	volatility := CalculateVolatility(returns)
	fmt.Printf("Volatility: %.2f%%\n", volatility*100)
}

package vix

import (
	"fmt"
	"math"
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

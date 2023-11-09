package vix

import (
	"github.com/theghostmac/volatilityindex/data"
	"math"
)

/*
Identify the options used in the calculation.
These are SPX put and call options with more than 23 days and less than 37 days to expiration.

Determine the precise time to expiration in minutes.

Identify the risk-free interest rate.
This is usually the Treasury Bond yield that more closely matches the option expiration date.

Calculate the contributions of individual options.

Calculate the volatility for near-term and next-term.

Finally, calculate the VIX index number.
*/

func CalculateVIX(options []data.Option, riskFreeRate float64, timeToExpiration float64) float64 {
	total := 0.0
	for _, option := range options {
		total += option.Premium * math.Exp(riskFreeRate*timeToExpiration) / math.Pow(option.StrikePrice, 2)
	}
	vix := math.Sqrt(total*2/timeToExpiration) * 100
	return vix
}

# Volatility Index

Also known as the Fear Gauge.

## Resources:
1. [CBOE Volatility Index (VIX)](https://markets.cboe.com/tradable_products/vix/)
2. [CBOE VIX White paper](https://www.math.wustl.edu/~victor/classes/ma456/vixwhite.pdf).
3. [Dummy data on SPX Puts and Calls](https://www.barchart.com/stocks/quotes/$SPX/options?expiration=2023-11-08-w)

## TO DO
- [x] Get S&P 500 Options data from Kaggle.
- [x] Write algorithm to calculate returns and volatility for all prices in the data.
- [x] Healthy refactor of code.
- [ ] Model the returned to be `PriceData` and have puts and calls
- [ ] Calculate expected volatility by aggregating weighted prices of SPX puts and calls over a range of strike prices. 
- [ ] Model data from external sources instead of CSV.
- [ ] Write a pipe for fetching data of different Option types.
- [ ] Pipe the data to the calculators.

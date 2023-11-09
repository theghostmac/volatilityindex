package utils

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

// ReadCSVFile reads the CSV file and returns the records excluding the header row.
func ReadCSVFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read the header row
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	// Read the remaining rows.
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

// ParseDateAndValue parses the date and value columns from CSV records.
func ParseDateAndValue(records [][]string) ([]time.Time, []float64, error) {
	var dates []time.Time
	var values []float64

	for _, record := range records {
		date, err := time.Parse("2006-01-02", record[0])
		if err != nil {
			return nil, nil, err
		}
		dates = append(dates, date)

		value, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, nil, err
		}
		values = append(values, value)
	}

	return dates, values, nil
}

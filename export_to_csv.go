package kekasigohelper

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ExportToCSV(outputFile string, headers []string, record [][]string) (err error) {
	// Open a new CSV file
	file, err := os.Create(outputFile)
	if err != nil {
		panic("Failed to create CSV file")
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header to the CSV file
	writer.Write(headers)

	// Write the record to the CSV file
	for _, v := range record {
		writer.Write(v)
	}

	// Flush any buffered data to the CSV file
	writer.Flush()

	// Check for errors during the write
	if err := writer.Error(); err != nil {
		return fmt.Errorf("CSV export failed : %v", err)
	}

	return nil
}

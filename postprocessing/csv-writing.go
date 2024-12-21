package postprocessing

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func saveResultToFile(data []map[string]string, fileName string) {
	// Get headers from the first map
	headers := getHeaders(data)

	// Create a CSV file
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write CSV data
	if err := writeCSV(file, headers, data); err != nil {
		fmt.Println("Error writing CSV:", err)
	}
}

// getHeaders extracts keys from the first map as headers
func getHeaders(data []map[string]string) []string {
	if len(data) == 0 {
		return nil
	}
	headers := []string{}
	for key := range data[0] {
		headers = append(headers, key)
	}
	return headers
}

// writeCSV writes the list of maps to a CSV file
func writeCSV(file io.Writer, headers []string, data []map[string]string) error {
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers
	if err := writer.Write(headers); err != nil {
		return err
	}

	// Write rows
	for _, row := range data {
		record := []string{}
		for _, header := range headers {
			record = append(record, row[header])
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

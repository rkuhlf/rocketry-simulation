package postprocessing

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func saveResultToFile(data []map[string]string, fileName string) error {
	// Get headers from the first map
	headers := getHeaders(data)

	dirPath := filepath.Dir(fileName)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	if err := writeCSV(file, headers, data); err != nil {
		return fmt.Errorf("error writing CSV %v", err)
	}

	return nil
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

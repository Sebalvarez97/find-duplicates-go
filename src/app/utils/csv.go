package utils

import (
	"encoding/csv"
	"os"
)

func ReadCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func SaveResultsToCSV(data [][]string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.WriteAll(data)
}

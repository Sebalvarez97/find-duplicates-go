package main

import (
	"find-duplicates/src/app/utils"

	"fmt"
	"log"
)

// Score scale
// More than 100 its High
// More than 50 its Medium
// Less than 35 its Low

func main() {
	route := "src/app/data/sample_data.csv"
	fmt.Println("Reading file from route:", route)

	data, err := utils.ReadCSV(route)
	if err != nil {
		log.Fatalf("Error reading CSV: %v", err)
	}

	var results [][]string
	results = append(results, []string{"ContactID Source", "ContactID Match", "Accuracy"})

	// Compare each row
	for i, row1 := range data {
		for j, row2 := range data {
			if i == j {
				continue
			}
			if score := utils.RowCompare(row1, row2); score > 35 {
				accuracy := "Low"
				if score > 50 {
					accuracy = "Medium"
				}
				if score > 100 {
					accuracy = "High"
				}
				results = append(results, []string{row1[0], row2[0], accuracy})
			}
		}
	}

	// Save results
	if err := utils.SaveResultsToCSV(results, "results.csv"); err != nil {
		log.Fatalf("Error saving results: %v", err)
	}

	fmt.Println("Results saved to results.csv")
}

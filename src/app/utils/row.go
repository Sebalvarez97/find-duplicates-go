package utils

import (
	"log"
	"strings"

	"github.com/xrash/smetrics"
)

// Define score weights
var columnScoreWeights = map[string]int{
	"contactID": 100,
	"name":      10,
	"name1":     10,
	"email":     50,
	"postalZip": 1,
	"address":   30,
}

func RowCompare(row1, row2 []string) int {
	if len(row1) != len(row2) {
		log.Panic("Row lengths do not match")
	}

	score := 0
	for i, value1 := range row1 {
		value2 := row2[i]
		column := GetColumnName(i)
		weight := columnScoreWeights[column]

		if strings.TrimSpace(value1) == strings.TrimSpace(value2) {
			score += weight
			continue
		}

		// Compare strings using JaroWinkler algorithm
		jaro := JaroMatch(value1, value2)

		switch column {
		case "contactID", "postalZip":
			// This collumn should be equal nor similar
			continue
		case "name", "name1":
			// Check initial letter
			if value1[0] == value2[0] {
				score += weight
			}
			continue
		case "email", "address":
			if jaro > 88 {
				difference := jaro - 88
				score += weight + difference
			}
			continue
		default:
			// Check similarity using JaroWinkler
			if jaro > 96 {
				score += weight
			}
		}

	}
	return score
}

func GetColumnName(index int) string {
	columns := []string{"contactID", "name", "name1", "email", "postalZip", "address"} // Adjust as per CSV
	if index < len(columns) {
		return columns[index]
	}
	return ""
}

// Using jarowinkler algorithm to compare strings
func JaroMatch(str1, str2 string) int {
	return int(smetrics.JaroWinkler(str1, str2, 0.7, 4) * 100)
}

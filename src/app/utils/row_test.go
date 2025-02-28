package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRowCompare_Different(t *testing.T) {
	row1 := []string{"7", "Gannon", "Weaver", "in.at@protonmail.org", "67238", "P.O. Box 410, 8379 Libero St."}
	row2 := []string{"8", "Hilary", "Franco", "congue.in@icloud.com", "25211", "331-981 Velit. Road"}
	score := RowCompare(row1, row2)
	assert.Equal(t, 0, score)
}

func TestRowCompare_SameName(t *testing.T) {
	row1 := []string{"7", "Gannon", "Weaver", "in.at@protonmail.org", "67238", "P.O. Box 410, 8379 Libero St."}
	row2 := []string{"8", "Gannon", "Franco", "congue.in@icloud.com", "25211", "331-981 Velit. Road"}
	score := RowCompare(row1, row2)
	assert.Equal(t, 10, score)
}

func TestRowCompare_SameName_And_Email(t *testing.T) {
	row1 := []string{"7", "Gannon", "Weaver", "in.at@protonmail.org", "67238", "P.O. Box 410, 8379 Libero St."}
	row2 := []string{"8", "Gannon", "Franco", "in.at@protonmail.org", "25211", "331-981 Velit. Road"}
	score := RowCompare(row1, row2)
	assert.Equal(t, 60, score)
}

func TestRowCompare_SameEmail_And_PostalCode(t *testing.T) {
	row1 := []string{"7", "Gannon", "Weaver", "in.at@protonmail.org", "67238", "P.O. Box 410, 8379 Libero St."}
	row2 := []string{"8", "G", "Franco", "in.at@protonmail.org", "67238", "331-981 Velit. Road"}
	score := RowCompare(row1, row2)
	assert.Equal(t, 65, score)
}

func TestRowCompare_DifferentLenghtError(t *testing.T) {
	row1 := []string{"7", "Gannon", "Weaver", "in.at@protonmail.org"}
	row2 := []string{"8", "Gannon", "Franco"}
	assert.Panics(t, func() { RowCompare(row1, row2) })
}

func TestRowCompare_SimilarAddress(t *testing.T) {
	row1 := []string{"7", "Gannon", "Weaver", "in.at@protonmail.org", "67238", "P.O. Box 410, 8379 Libero St."}
	row2 := []string{"8", "Hilary", "Franco", "congue.in@icloud.com", "25211", "P.O. Box 410, 8872 Libero Street"}
	score := RowCompare(row1, row2)
	assert.Equal(t, 30, score)
}

func TestRowCompare_SimilarAddress2(t *testing.T) {
	row1 := []string{"17", "Vance", "Martin", "rat.nonummy@icloud.net", "78358", "694-3213 Diam St."}
	row2 := []string{"517", "V", "Martin", "erat.nonummy@icloud.net", "78358", "694-3213 Diam St."}
	score := RowCompare(row1, row2)
	assert.Equal(t, 105, score)
}

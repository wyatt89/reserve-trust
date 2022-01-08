package main

import (
	"testing"
)

func TestValid(t *testing.T) {

	tt := loadTestTable(t)

	for _, tc := range tt {
		r := readFromFile(t, tc.filepath)

		_, _ = ParsePayment(r)
	}

}

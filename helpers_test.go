package main

import (
	"os"
	"testing"
)

type paymentHelper struct {
	filepath      string
	shouldError   bool
	expectedError string
}

func loadTestTable(t *testing.T) []paymentHelper {
	t.Helper()

	return []paymentHelper{
		{
			filepath:    "./testdata/valid0.txt",
			shouldError: false,
		},
		{
			filepath:    "./testdata/valid1.txt",
			shouldError: false,
		},
		{
			filepath:    "./testdata/valid2.txt",
			shouldError: false,
		},
		{
			filepath:    "./testdata/a0.txt",
			shouldError: true,
		},
		{
			filepath:    "./testdata/a1.txt",
			shouldError: true,
		},
		{
			filepath:    "./testdata/b0.txt",
			shouldError: true,
		},
		{
			filepath:    "./testdata/b1.txt",
			shouldError: true,
		},
		{
			filepath:    "./testdata/c0.txt",
			shouldError: true,
		},
		{
			filepath:    "./testdata/c1.txt",
			shouldError: true,
		},
		{
			filepath:    "./testdata/d0.txt",
			shouldError: true,
		},
		{
			filepath:    "./testdata/d1.txt",
			shouldError: true,
		},
		{
			filepath:    "./testdata/e0.txt",
			shouldError: true,
		},
		{
			filepath:    "./testdata/e1.txt",
			shouldError: true,
		},
		{
			filepath:    "./testdata/f0.txt",
			shouldError: true,
		},
		{
			filepath:    "./testdata/f1.txt",
			shouldError: true,
		},
	}

}

func readFromFile(t *testing.T, filepath string) []byte {
	t.Helper()

	dat, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	return dat
}

package main

import (
	"os"
	"testing"
)

type paymentHelper struct {
	filepath      string
	shouldError   bool
	expectedError string
	Payment
}

func loadTestTable(t *testing.T) []paymentHelper {
	t.Helper()

	return []paymentHelper{
		{
			filepath:    "./testdata/valid0.txt",
			shouldError: false,
			Payment:     Payment{"Clark Kent", 123123123, 678678678, "thanks for dinner", 345345345, 9090909},
		},
		{
			filepath:    "./testdata/valid1.txt",
			shouldError: false,
			Payment:     Payment{"Peter Parker", 12345671, 678637892, "breakfast was great", 345345345, 9090909},
		},
		{
			filepath:    "./testdata/valid2.txt",
			shouldError: false,
			Payment:     Payment{"Jefferson Pierce", 9845671, 67234678, "", 349180345, 829389729},
		},
		{
			filepath:      "./testdata/a0.txt",
			shouldError:   true,
			expectedError: "error: SenderName is invalid",
		},
		{
			filepath:      "./testdata/a1.txt",
			shouldError:   true,
			expectedError: "error: SenderName is invalid",
		},
		{
			filepath:      "./testdata/b0.txt",
			shouldError:   true,
			expectedError: "error: SenderAccountNumber is invalid",
		},
		{
			filepath:      "./testdata/b1.txt",
			shouldError:   true,
			expectedError: "error: SenderAccountNumber is invalid",
		},
		{
			filepath:      "./testdata/c0.txt",
			shouldError:   true,
			expectedError: "error: SenderRoutingNumber is invalid",
		},
		{
			filepath:      "./testdata/c1.txt",
			shouldError:   true,
			expectedError: "error: SenderRoutingNumber is invalid",
		},
		{
			filepath:      "./testdata/d0.txt",
			shouldError:   true,
			expectedError: "error: payment has too few or too many values",
		},
		{ // opted to split string all at once instead parsing section by section. This made knowing which section was invalid difficult
			filepath:      "./testdata/d1.txt",
			shouldError:   true,
			expectedError: "error: SenderNote is invalid",
		},
		{
			filepath:      "./testdata/e0.txt",
			shouldError:   true,
			expectedError: "error: RecipientAccountNumber is invalid",
		},
		{
			filepath:      "./testdata/e1.txt",
			shouldError:   true,
			expectedError: "error: RecipientAccountNumber is invalid",
		},
		{
			filepath:      "./testdata/f0.txt",
			shouldError:   true,
			expectedError: "error: RecipientRoutingNumber is invalid",
		},
		{
			filepath:      "./testdata/f1.txt",
			shouldError:   true,
			expectedError: "error: RecipientRoutingNumber is invalid",
		},
		{
			filepath:      "./testdata/g4-sameroutingandaccountnum.txt",
			shouldError:   true,
			expectedError: "error: send account and routing number cannot match receiver account and routing number",
		},
		{
			filepath:      "./testdata/g1-emptycontent.txt",
			shouldError:   true,
			expectedError: "error: payment has too few or too many values",
		},
		{
			filepath:      "./testdata/g2-emptyfile.txt",
			shouldError:   true,
			expectedError: "error: payment file is empty",
		},
		{
			filepath:      "./testdata/g0-coloninsenderaccountnumber.txt",
			shouldError:   true,
			expectedError: "error: payment has too few or too many values",
		},
		{
			filepath:      "./testdata/g5-toofewfields.txt",
			shouldError:   true,
			expectedError: "error: payment has too few or too many values",
		},
		{
			filepath:      "./testdata/g3-invalidstartsymbol.txt",
			shouldError:   true,
			expectedError: "error: start or end symbol is not present",
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

func loadMultiPaymentTestTable(t *testing.T) []paymentHelper {
	t.Helper()

	return []paymentHelper{
		{
			filepath:    "./testdata/multipayment/valid3.txt",
			shouldError: false,
		},
	}

}

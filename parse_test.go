package main

import (
	"testing"
)

func TestValid(t *testing.T) {

	tt := loadTestTable(t)

	for _, tc := range tt {
		r := readFromFile(t, tc.filepath)

		actP, err := ParsePayment(r)

		if tc.shouldError && tc.expectedError != err.Error() {
			t.Fail()
		}
		if !tc.shouldError {
			if actP.SenderName != tc.Payment.SenderName {
				t.Fail()
			}
			if actP.RecipientRoutingNumber != tc.Payment.RecipientRoutingNumber {
				t.Fail()
			}
			// can easily add checks on all payment fields if required
		}

	}

}

//
//func TestValidMultiPaymentFile(t *testing.T) {
//
//	tt := loadMultiPaymentTestTable(t)
//
//	for _, tc := range tt {
//		r := readFromFile(t, tc.filepath)
//
//		fmt.Println(tc.filepath)
//		_, _ = ParsePayments(r)
//		fmt.Println("---------------------------------")
//	}
//
//}

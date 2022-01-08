package main

import (
	"errors"
	"strings"
)

type Payment struct {
	SenderName             string
	SenderAccountNumber    int
	SenderRoutingNumber    int
	SenderNote             string
	RecipientAccountNumber int
	RecipientRoutingNumber int
}

var (
	ErrorInvalidSenderName             = errors.New("error: SenderName is invalid")
	ErrorInvalidSenderAccountNumber    = errors.New("error: SenderAccountNumber is invalid")
	ErrorInvalidSenderRoutingNumber    = errors.New("error: SenderRoutingNumber is invalid")
	ErrorInvalidSenderNote             = errors.New("error: SenderNote is invalid")
	ErrorInvalidRecipientAccountNumber = errors.New("error: RecipientAccountNumber is invalid")
	ErrorInvalidRecipientRoutingNumber = errors.New("error: RecipientRoutingNumber is invalid")
)

func ParsePayment(payment []byte) (Payment, error) {
	return Payment{}, nil
}

func ParseMultiLineFile(str, before, after string) (string, string) {
	a := strings.SplitAfterN(str, before, 2)
	b := strings.SplitAfterN(a[len(a)-1], after, 2)
	if 1 == len(b) {
		return b[0], ""
	}
	return b[0][0 : len(b[0])-len(after)], b[1]
}

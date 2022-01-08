package main

import "errors"

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

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Payment Leading zeroes are valid in account and routing numbers, I think these fields are probably better represented as strings.
type Payment struct {
	SenderName             string
	SenderAccountNumber    int
	SenderRoutingNumber    int
	SenderNote             string
	RecipientAccountNumber int
	RecipientRoutingNumber int
}

const (
	ExpectedValueCount = 6
	MaxCharLength      = 100
	MaxIntLength       = 100
)

const (
	SenderNameInd = iota
	SenderAccountNumberInd
	SenderRoutingNumberInd
	SenderNoteInd
	RecipientAccountNumberInd
	RecipientRoutingNumberInd
)

var (
	ErrorInvalidSenderName             = errors.New("error: SenderName is invalid")
	ErrorInvalidSenderAccountNumber    = errors.New("error: SenderAccountNumber is invalid")
	ErrorInvalidSenderRoutingNumber    = errors.New("error: SenderRoutingNumber is invalid")
	ErrorInvalidSenderNote             = errors.New("error: SenderNote is invalid")
	ErrorInvalidRecipientAccountNumber = errors.New("error: RecipientAccountNumber is invalid")
	ErrorInvalidRecipientRoutingNumber = errors.New("error: RecipientRoutingNumber is invalid")
	ErrorInvalidPaymentFile            = errors.New("error: payment file is empty")

	// ErrorInvalidPaymentFormat opted to split payment string all at once instead of parsing field by field. This made identifying which section was invalid difficult
	// could change it to parse field by field
	// current implementation will catch if the note or any other field contains a colon, because it is a file the integer values could contain ":"
	ErrorInvalidPaymentFormat                             = errors.New("error: payment has too few or too many values")
	ErrorMatchingSenderAndReceiverAccountAndRoutingNumber = errors.New("error: send account and routing number cannot match receiver account and routing number")
	ErrorInvalidSurroundingSymbols                        = errors.New("error: start or end symbol is not present")
)

func ParsePayment(paymentData []byte) (Payment, error) {
	if len(paymentData) == 0 { // Validate file is not empty
		return Payment{}, ErrorInvalidPaymentFile
	}

	res, _, err := parsePaymentFromFile(string(paymentData), "{", "}")
	if err != nil { // Validate start and ending symbol is present for single payment parser
		return Payment{}, err
	}

	paymentFields := strings.Split(res, ":")
	if len(paymentFields) != ExpectedValueCount { // Validate that only the 6 expected fields are present
		return Payment{}, ErrorInvalidPaymentFormat
	}

	p := Payment{}
	//SenderName
	s, err := validatePaymentString(paymentFields, SenderNameInd, ErrorInvalidSenderName)
	if err != nil {
		return p, err
	}
	p.SenderName = s

	//SenderAccountNumber
	i, err := validatePaymentInt(paymentFields, SenderAccountNumberInd, ErrorInvalidSenderAccountNumber)
	if err != nil {
		return p, err
	}
	p.SenderAccountNumber = i

	//SenderRoutingNumber
	i, err = validatePaymentInt(paymentFields, SenderRoutingNumberInd, ErrorInvalidSenderRoutingNumber)
	if err != nil {
		return p, err
	}
	p.SenderRoutingNumber = i

	//SenderNote
	s, err = validatePaymentString(paymentFields, SenderNoteInd, ErrorInvalidSenderNote)
	if err != nil {
		return p, err
	}
	p.SenderNote = s

	//RecipientAccountNumber
	i, err = validatePaymentInt(paymentFields, RecipientAccountNumberInd, ErrorInvalidRecipientAccountNumber)
	if err != nil {
		return p, err
	}
	p.RecipientAccountNumber = i

	//RecipientRoutingNumber
	i, err = validatePaymentInt(paymentFields, RecipientRoutingNumberInd, ErrorInvalidRecipientRoutingNumber)
	if err != nil {
		return p, err
	}
	p.RecipientRoutingNumber = i

	// Validate the Sender and Receiver do not have matching account and routing number as a combination
	senderAccRout := fmt.Sprintf("%v%v", p.SenderAccountNumber, p.SenderRoutingNumber)
	recAccRout := fmt.Sprintf("%v%v", p.RecipientAccountNumber, p.RecipientRoutingNumber)

	if senderAccRout == recAccRout {
		return p, ErrorMatchingSenderAndReceiverAccountAndRoutingNumber
	}

	return p, nil
}

// Validate Max character length for SenderName and SenderNote - Name cannot be empty
func validatePaymentString(paymentFields []string, pInd int, err error) (string, error) {
	str := paymentFields[pInd]

	if len(str) > MaxCharLength {
		return "-1", err
	}
	if pInd == SenderNameInd && len(str) == 0 {
		return "-1", err
	}
	return str, nil
}

// Validate the max integer length and that the int is parseable
func validatePaymentInt(paymentFields []string, pInd int, err error) (int, error) {
	str := paymentFields[pInd]

	if len(str) > MaxIntLength || str == "" {
		return -1, err
	}

	// Could add an extra validation that the numeric string only contains numbers then wrap it in the field error that is passed in
	num, conErr := strconv.Atoi(paymentFields[pInd])
	if conErr != nil {
		return -1, err
	}

	return num, nil
}

// Can easily parse multiple payments in a file by being called in a loop and passing the remainder b[1] back to parsePaymentFromFile()
// Can change surrounding start and end symbols from "{", "}" to another arbitrary symbol
func parsePaymentFromFile(pData, before, after string) (string, string, error) {
	validStart := strings.Contains(pData, "{")
	validEnd := strings.Contains(pData, "}")
	if !validEnd || !validStart {
		return "", "", ErrorInvalidSurroundingSymbols
	}
	a := strings.SplitAfterN(pData, before, 2)
	b := strings.SplitAfterN(a[len(a)-1], after, 2)
	if 1 == len(b) {
		return b[0], "", nil
	}
	res := strings.ReplaceAll(b[0][0:len(b[0])-len(after)], "\n", "")
	return res, b[1], nil
}

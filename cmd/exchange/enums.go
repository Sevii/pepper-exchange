package main

import (
	"bytes"
	"encoding/json"
)

type Exchange int
type Operation int

const (
	ASK    Operation = iota // value: 1, type: BookOperation
	BID                     // value: 2, type: BookOperation
	CANCEL                  // value: 3, type: BookOperation
	STATUS
	INVALID_OPERATION
)

const (
	BTCUSD  Exchange = iota // value: 1, type: Exchange
	BTCLTC                  // value: 2, type: Exchange
	BTCDOGE                 // value: 3, type: Exchange
	BTCXMR                  // value: 4, type: Exchange
	INVALID_EXCHANGE
)

func (exchange Exchange) String() string {
	// declare an array of strings in the same order as the Exchange enum
	names := [...]string{
		"BTCUSD",
		"BTCLTC",
		"BTCDOGE",
		"BTCXMR",
		"INVALID_EXCHANGE"}

	// Prevent panicking in case exchange  is out of range of the enum
	if exchange < BTCUSD || exchange > INVALID_EXCHANGE {
		return "Unknown"
	}
	// Returns the
	return names[exchange]
}

func ExchangeFromStr(str string) Exchange {
	switch str {
	case "BTCUSD":
		return BTCUSD
	case "BTCLTC":
		return BTCLTC
	case "BTCDOGE":
		return BTCDOGE
	case "BTCXMR":
		return BTCXMR
	default:
		return INVALID_EXCHANGE

	}

}

func (operation Operation) String() string {
	// declare an array of strings in the same order as the Exchange enum
	names := [...]string{
		"ASK",
		"BID",
		"CANCEL",
		"STATUS",
		"INVALID"}

	// Prevent panicking in case exchange  is out of range of the enum
	if operation < ASK || operation > INVALID_OPERATION {
		return "Invalid"
	}
	// Returns the
	return names[operation]
}

func OperationFromStr(str string) Operation {
	switch str {
	case "ask":
		return ASK
	case "bid":
		return BID
	case "cancel":
		return CANCEL
	case "status":
		return STATUS
	case "ASK":
		return ASK
	case "BID":
		return BID
	case "CANCEL":
		return CANCEL
	case "STATUS":
		return STATUS
	default:
		return INVALID_OPERATION
	}

}

func (d *Exchange) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(d.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (d *Exchange) UnmarshalJSON(b []byte) error {
	// unmarshal as string
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	// lookup value
	*d = ExchangeFromStr(s)
	return nil
}

func (d *Operation) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(d.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (d *Operation) UnmarshalJSON(b []byte) error {
	// unmarshal as string
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	// lookup value
	*d = OperationFromStr(s)
	return nil
}

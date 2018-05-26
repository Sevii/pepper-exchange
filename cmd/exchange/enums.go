package main

import (
	"fmt"
)

type Exchange int
type Operation int

const (
	ASK    Operation = iota // value: 1, type: BookOperation
	BID                     // value: 2, type: BookOperation
	CANCEL                  // value: 3, type: BookOperation
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
	fmt.Println("BTCLTC", BTCLTC.String())
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
		"INVALID"}

	// Prevent panicking in case exchange  is out of range of the enum
	if operation < ASK || operation > INVALID_OPERATION {
		return "Invalid"
	}
	// Returns the
	return names[operation]
}

func OperationFromStr(str string) Operation {
	fmt.Println(str)
	switch str {
	case "ask":
		fmt.Println(str, ASK)
		return ASK
	case "bid":
		fmt.Println(str, BID)
		return BID
	case "cancel":
		fmt.Println(str, CANCEL)
		return CANCEL
	default:
		return INVALID_OPERATION
	}

}

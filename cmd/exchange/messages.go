package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"time"
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

type BookOperation interface {
	Type() Operation
}

// {"id": 123, "direction": "bid", "exchange":"BTCUSD", "number":123,"price":1000 }
//Order is any bid or ask on the exchange
type Order struct {
	ID        uuid.UUID // The id of the order
	Direction Operation // Whether this order is buying (bid) or selling (ask)
	Exchange  Exchange  // The exchange either BTC/USD, BTC/LTC, BTC/Doge, BTC/XMR(Monero)
	Number    int       // The number of coins
	Price     int       //price is always in Satoshis
	Timestamp int       // timestamp in nanoseconds
}

func NewOrder(req OrderRequest) Order {
	uid := uuid.NewV4()
	fmt.Println("req.Direction", req.Direction)
	return Order{
		ID:        uid,
		Direction: OperationFromStr(req.Direction),
		Exchange:  ExchangeFromStr(req.Exchange),
		Number:    req.Number,
		Price:     req.Price,
		Timestamp: time.Now().Nanosecond()}
}

//Cancel is a request to cancel an outstanding order. Only non-filled parts of an order can be canceled.
type Cancel struct {
	ID        uuid.UUID
	Order_id  uuid.UUID
	Exchange  Exchange
	Timestamp int
}

func NewCancel(req CancelRequest) Cancel {
	uid := uuid.NewV4()

	return Cancel{
		ID:        uid,
		Order_id:  req.Order_id,
		Exchange:  req.Exchange,
		Timestamp: time.Now().Nanosecond()}
}

//Fill is a match between a bid and ask for x satoshis and y number of coins
type Fill struct {
	ID        uuid.UUID `json:"id"`
	Exchange  Exchange  `json:"exchange"`
	Number    int       `json:"number"`
	Price     int       `json:"price"`
	Ask_id    uuid.UUID `json:"ask_id"`
	Bid_id    uuid.UUID `json:"bid_id"`
	Timestamp int       `json:"timestamp"`
}

func NewFill(exc Exchange, num int, price int, ask_id uuid.UUID, bid_id uuid.UUID) Fill {
	uid := uuid.NewV4()

	return Fill{
		ID:        uid,
		Exchange:  exc,
		Number:    num,
		Price:     price,
		Ask_id:    ask_id,
		Bid_id:    bid_id,
		Timestamp: time.Now().Nanosecond()}
}

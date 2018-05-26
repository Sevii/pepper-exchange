package main

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"time"
)

type BookOperation interface {
	Type() Operation
}

// {"id": 123, "direction": "bid", "exchange":"BTCUSD", "number":123,"price":1000 }
//Order is any bid or ask on the exchange
type Order struct {
	ID                uuid.UUID // The id of the order
	Direction         Operation // Whether this order is buying (bid) or selling (ask)
	Exchange          Exchange  // The exchange either BTC/USD, BTC/LTC, BTC/Doge, BTC/XMR(Monero)
	Number            int       // The number of coins
	NumberOutstanding int       //Outstanding coins
	Price             int       //price is always in Satoshis
	Timestamp         int       // timestamp in nanoseconds
}

func NewOrder(req OrderRequest) Order {
	uid := uuid.NewV4()
	return Order{
		ID:                uid,
		Direction:         OperationFromStr(req.Direction),
		Exchange:          ExchangeFromStr(req.Exchange),
		Number:            req.Number,
		NumberOutstanding: req.Number,
		Price:             req.Price,
		Timestamp:         time.Now().Nanosecond()}
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
		Exchange:  ExchangeFromStr(req.Exchange),
		Timestamp: time.Now().Nanosecond()}
}

//Fill is a match between a bid and ask for x satoshis and y number of coins
type Fill struct {
	ID           uuid.UUID `json:"id"`
	Exchange     Exchange  `json:"exchange"`
	Number       int       `json:"number"`
	Price        int       `json:"price"`
	Timestamp    int       `json:"timestamp"`
	Participants []Order   `json:"participants"`
	Closed       []Order   `json:"closed"`
}

func NewFill(exc Exchange, num int, price int, part []Order, closed []Order) Fill {
	uid := uuid.NewV4()

	return Fill{
		ID:           uid,
		Exchange:     exc,
		Number:       num,
		Price:        price,
		Timestamp:    time.Now().Nanosecond(),
		Participants: part,
		Closed:       closed}
}

func String(s Fill) string {
	out, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	return fmt.Sprint(string(out))
}

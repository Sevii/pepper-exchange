package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

// {"id": 123, "direction": "bid", "exchange":"btcUsd", "number":123,"price":1000 }
//OrderRequest struct used to submit an ask or bid to the exchange
type OrderRequest struct {
	Direction string `json:"direction"` // Whether this order is buying (bid) or selling (ask)
	Exchange  string `json:"exchange"`  // The exchange either BTC/USD, BTC/LTC, BTC/Doge, BTC/XMR(Monero)
	Number    int    `json:"number"`    // The number of coins
	Price     int    `json:"price"`     //price is always in Satoshis
	UserID    string `json:"userId"`    // User id
}

type CancelRequest struct {
	OrderID  uuid.UUID `json:"orderId"`
	Exchange string    `json:"exchange"`
	UserID   string    `json:"userId"` // User id
}

var netClient = &http.Client{
	Timeout: time.Second * 3,
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	var ord OrderRequest

	//Check for Authorization header
	// if r.Host

	// Deserialize the order
	err := json.NewDecoder(r.Body).Decode(&ord)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//Validate required fields are present

	//Validate the User has enough coins to make the trade

	//Create order struct and timestamp it
	order := NewOrder(ord)
	switch order.Exchange {
	case BTCUSD:
		toOrderBooks[BTCUSD] <- order
	case BTCLTC:
		toOrderBooks[BTCLTC] <- order
	case BTCDOGE:
		toOrderBooks[BTCDOGE] <- order
	case BTCXMR:
		toOrderBooks[BTCXMR] <- order
	default:
		http.Error(w, "Nonexistent Exchange requested", 400)
		return
	}
	//Send Order to OrderBook chan

	//Update Redis with the order

	//Return 200
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)

}

type StatusRequest struct {
	Exchange string `json:"exchange"`
	UserID   string `json:"userId"` // User id
}

func orderStatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	exchange := vars["exchange"]
	exe := ExchangeFromStr(exchange)
	if exe == INVALID_EXCHANGE {
		http.Error(w, "Nonexistent Exchange requested", 400)
		return
	}
	userId := vars["userId"]
	statusOrder := NewStatusOrder(StatusRequest{UserID: userId, Exchange: exchange})

	var orders []Order
	//Validate required fields are present

	//Create Cancel struct and timestamp it

	//Create order struct and timestamp it

	switch statusOrder.Exchange {
	case BTCUSD:
		toOrderBooks[BTCUSD] <- statusOrder
		orders = <-fromOrderBooks[BTCUSD]

	case BTCLTC:
		toOrderBooks[BTCLTC] <- statusOrder
		orders = <-fromOrderBooks[BTCLTC]
	case BTCDOGE:
		toOrderBooks[BTCDOGE] <- statusOrder
		orders = <-fromOrderBooks[BTCDOGE]
	case BTCXMR:
		toOrderBooks[BTCXMR] <- statusOrder
		orders = <-fromOrderBooks[BTCXMR]
	default:
		http.Error(w, "Nonexistent Exchange requested", 400)
		return
	}

	//Send Cancel to OrderBook chan

	//Update Redis with the cancelation

	//Return 200
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(orders)
}

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	var cancel CancelRequest

	// Deserialize the order
	err := json.NewDecoder(r.Body).Decode(&cancel)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	//Validate required fields are present

	//Create Cancel struct and timestamp it

	//Create order struct and timestamp it
	cancelOrder := NewCancelOrder(cancel)
	switch cancelOrder.Exchange {
	case BTCUSD:
		toOrderBooks[BTCUSD] <- cancelOrder
	case BTCLTC:
		toOrderBooks[BTCLTC] <- cancelOrder
	case BTCDOGE:
		toOrderBooks[BTCDOGE] <- cancelOrder
	case BTCXMR:
		toOrderBooks[BTCXMR] <- cancelOrder
	default:
		http.Error(w, "Nonexistent Exchange requested", 400)
		return
	}

	//Send Cancel to OrderBook chan

	//Update Redis with the cancelation

	//Return 200
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	w.Write([]byte("{\"alive\": true}"))
	// io.WriteString(w, `{"alive": true}`)
}

package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"log"
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

	if ord.Price < 0 || ord.Number <= 0 {
		http.Error(w, "Invalid order parameters, number of coins or price is unacceptable.", 400)
		return
	}

	// Create Order struct
	order := NewOrder(ord)

	//Validate the User has enough coins to make the trade
	valid := ValidateOrder(order)
	if valid != true {
		http.Error(w, "You cannot afford it.", 400)
		return
	}

	//Create order struct and timestamp it

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

func marketDataHandler(w http.ResponseWriter, r *http.Request) {
	market, err := getMarketData()
	if err != nil {
		http.Error(w, "Failed to get Market data", 500)
		log.Println(err)
		return
	}

	//Return 200
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(market)
}

func accountStatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId := vars["userId"]
	accountStatus, err := getAccountStatusRedis(userId)
	if err != nil {
		http.Error(w, "Failed to get your account", 500)
		return
	}

	//Return 200
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(accountStatus)
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

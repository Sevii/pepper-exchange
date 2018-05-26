package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

var (
	toOrderBooks map[Exchange]chan Order
)

// {"id": 123, "direction": "bid", "exchange":"btcUsd", "number":123,"price":1000 }
//OrderRequest struct used to submit an ask or bid to the exchange
type OrderRequest struct {
	Direction string `json:"direction"` // Whether this order is buying (bid) or selling (ask)
	Exchange  string `json:"exchange"`  // The exchange either BTC/USD, BTC/LTC, BTC/Doge, BTC/XMR(Monero)
	Number    int    `json:"number"`    // The number of coins
	Price     int    `json:"price"`     //price is always in Satoshis
}

type CancelRequest struct {
	Order_id uuid.UUID `json:"order_id"`
	Exchange Exchange  `json:"exchange"`
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
	fmt.Println("order: ", order)
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

//setupOrderBooks populates the toOrderBooks global map
func setupOrderBooks() {

	toOrderBooks = make(map[Exchange]chan Order)

	btcUsd := NewBookManager(BTCUSD)
	btcUsdIn := make(chan Order)
	btcUsdOut := make(chan Fill)

	toOrderBooks[btcUsd.Exchange] = btcUsdIn
	go btcUsd.Run(btcUsdIn, btcUsdOut)

	btcLtc := NewBookManager(BTCLTC)
	btcLtcIn := make(chan Order)
	btcLtcOut := make(chan Fill)

	toOrderBooks[btcLtc.Exchange] = btcLtcIn
	go btcUsd.Run(btcLtcIn, btcLtcOut)

	btcDoge := NewBookManager(BTCDOGE)
	btcDogeIn := make(chan Order)
	btcDogeOut := make(chan Fill)

	toOrderBooks[btcDoge.Exchange] = btcDogeIn
	go btcUsd.Run(btcDogeIn, btcDogeOut)

	btcXmr := NewBookManager(BTCXMR)
	btcXmrIn := make(chan Order)
	btcXmrOut := make(chan Fill)

	toOrderBooks[btcXmr.Exchange] = btcXmrIn
	go btcUsd.Run(btcXmrIn, btcXmrOut)
	fmt.Println("Finished starting channels")

}

func main() {

	setupOrderBooks()

	r := mux.NewRouter()

	r.HandleFunc("/health", HealthCheckHandler)
	r.HandleFunc("/order", orderHandler).Methods("POST")
	r.HandleFunc("/cancel", cancelHandler).Methods("POST")
	http.ListenAndServe(":8080", r)
}

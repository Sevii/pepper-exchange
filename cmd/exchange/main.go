package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

type Exchange int

const (
	BTCUSD  Exchange = iota + 1 // value: 1, type: Exchange
	BTCLTC                      // value: 2, type: Exchange
	BTCDOGE                     // value: 3, type: Exchange
	BTCXMR                      // value: 4, type: Exchange
	INVALID_EXCHANGE
)

var (
	toOrderBook chan Order
)

func (exchange Exchange) String() string {
	// declare an array of strings in the same order as the Exchange enum
	names := [...]string{
		"BTCUSD",
		"BTCLTC",
		"BTCDOGE",
		"BTCXMR"}

	// Prevent panicking in case exchange  is out of range of the enum
	if exchange < BTCUSD || exchange > BTCXMR {
		return "Unknown"
	}
	// Returns the
	return names[exchange]
}

func ExchangeFromStr(str string) Exchange {
	fmt.Println(str)
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

// {"id": 123, "direction": "bid", "exchange":"BTCUSD", "number":123,"price":1000 }
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
	fmt.Println(ord)
	//Validate required fields are present

	//Validate the User has enough coins to make the trade

	//Create order struct and timestamp it
	order := NewOrder(ord)
	fmt.Println("order: ", order)
	toOrderBook <- order
	//Send Order to OrderBook chan

	//Update Redis with the order

	//Return 200
	w.WriteHeader(http.StatusOK)

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

func main() {

	manager := NewBookManager()
	toOrderBook = make(chan Order)
	out := make(chan Fill)

	go manager.Run(toOrderBook, out)

	r := mux.NewRouter()

	r.HandleFunc("/health", HealthCheckHandler)
	r.HandleFunc("/order", orderHandler).Methods("POST")
	r.HandleFunc("/cancel", orderHandler).Methods("POST")
	http.ListenAndServe(":8080", r)
}

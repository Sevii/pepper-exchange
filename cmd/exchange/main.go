package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Exchange int

const (
	BTCUSD  Exchange = iota + 1 // value: 1, type: Exchange
	BTCLTC                      // value: 2, type: Exchange
	BTCDoge                     // value: 3, type: Exchange
	BTCXMR                      // value: 4, type: Exchange
)

func (exchange Exchange) String() string {
	// declare an array of strings in the same order as the Exchange enum
	names := [...]string{
		"BTCUSD",
		"BTCLTC",
		"BTCDoge",
		"BTCXMR"}

	// Prevent panicking in case exchange  is out of range of the enum
	if exchange < BTCUSD || exchange > BTCXMR {
		return "Unknown"
	}
	// Returns the
	return names[exchange]
}

//Order is any bid or ask on the exchange
type Order struct {
	id        int      // The id of the order
	direction string   // Whether this order is buying (bid) or selling (ask)
	exchange  Exchange // The exchange either BTC/USD, BTC/LTC, BTC/Doge, BTC/XMR(Monero)
	number    int      // The number of coins
	price     int      //price is always in Satoshis
	timestamp int      // timestamp in nanoseconds
}

//Fill is a match between a bid and ask for x satoshis and y number of coins
type Fill struct {
	id        int
	exchange  Exchange
	number    int
	price     int
	ask_id    int
	bid_id    int
	timestamp int
}

type Cancel struct {
	id        int
	exchange  Exchange
	order_id  int
	timestamp int
}

var netClient = &http.Client{
	Timeout: time.Second * 3,
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	var ord Order
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	//Check for Authorization header

	// Deserialize the order
	err := json.NewDecoder(r.Body).Decode(&ord)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	//Validate required fields are present

	//Validate the User has enough coins to make the trade

	//Create order struct and timestamp it

	//Send Order to OrderBook chan

	//Update Redis with the order

	//Return 200
	w.WriteHeader(http.StatusOK)

}

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	var cancel Cancel
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

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

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "REPORTS")
	var u user
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	weatherData, err := getWeather()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	s, err := getStocks()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	today := fmt.Sprintf(time.Now().Format(time.RFC850))

	report := Report{Username: u.Username, Level: u.Level, Dow: s.Dow, SP500: s.SP500, Temp: weatherData.Temp, Weather: weatherData.Type, Time: today}

	fmt.Println("Created report: ", report)

	w.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.Marshal(report)
	w.Write(jsonData)

}

func main() {
	http.HandleFunc("/report", handler)
	http.ListenAndServe(":8080", nil)
}

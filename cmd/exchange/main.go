package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Exchange int

const (
	BTCUSD  Exchange = iota + 1 // value: 1, type: Weekday
	BTCLTC                      // value: 2, type: Weekday
	BTCDoge                     // value: 3, type: Weekday
	BTCXMR                      // value: 4, type: Weekday
)

type Order struct {
	id        int      // The id of the order
	direction string   // Whether this order is buying (bid) or selling (ask)
	exchange  Exchange // The exchange either BTC/USD, BTC/LTC, BTC/Doge, BTC/XMR(Monero)
	number    int      // The number of coins
	price     int      //price is always in Satoshis
}

type Fill struct {
	id        int
	exchange  Exchange
	number    int
	price     int
	ask_id    int
	bid_id    int
	timestamp int
}

var netClient = &http.Client{
	Timeout: time.Second * 3,
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

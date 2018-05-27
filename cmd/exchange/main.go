package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	toOrderBooks   map[Exchange]chan Order
	fromOrderBooks map[Exchange]chan []Order
)

//setupOrderBooks populates the toOrderBooks global map
func setupOrderBooks() {

	toOrderBooks = make(map[Exchange]chan Order)
	fromOrderBooks = make(map[Exchange]chan []Order)

	btcUsd := NewBookManager(BTCUSD)
	btcUsdIn := make(chan Order)
	btcUsdOut := make(chan []Order)

	toOrderBooks[btcUsd.Exchange] = btcUsdIn
	fromOrderBooks[btcUsd.Exchange] = btcUsdOut

	go btcUsd.Run(btcUsdIn, btcUsdOut)

	btcLtc := NewBookManager(BTCLTC)
	btcLtcIn := make(chan Order)
	btcLtcOut := make(chan []Order)

	toOrderBooks[btcLtc.Exchange] = btcLtcIn
	fromOrderBooks[btcLtc.Exchange] = btcLtcOut

	go btcLtc.Run(btcLtcIn, btcLtcOut)

	btcDoge := NewBookManager(BTCDOGE)
	btcDogeIn := make(chan Order)
	btcDogeOut := make(chan []Order)

	toOrderBooks[btcDoge.Exchange] = btcDogeIn
	fromOrderBooks[btcDoge.Exchange] = btcDogeOut

	go btcDoge.Run(btcDogeIn, btcDogeOut)

	btcXmr := NewBookManager(BTCXMR)
	btcXmrIn := make(chan Order)
	btcXmrOut := make(chan []Order)

	toOrderBooks[btcXmr.Exchange] = btcXmrIn
	fromOrderBooks[btcXmr.Exchange] = btcXmrOut

	go btcXmr.Run(btcXmrIn, btcXmrOut)
	fmt.Println("Finished starting channels")

}

func main() {

	setupOrderBooks()

	r := mux.NewRouter()

	r.HandleFunc("/health", HealthCheckHandler)
	r.HandleFunc("/order", orderHandler).Methods("POST")
	r.HandleFunc("/cancel", cancelHandler).Methods("POST")
	r.HandleFunc("/order-status/{userId}/{exchange}", orderStatusHandler).Methods("GET")
	r.HandleFunc("/stream/fills/{exchange}", FillsWebsocketHandler).Methods("GET")
	http.ListenAndServe(":8080", r)
}

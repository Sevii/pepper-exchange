package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	toOrderBooks map[Exchange]chan Order
)

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

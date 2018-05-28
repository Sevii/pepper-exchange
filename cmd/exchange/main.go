package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "net/http/pprof"
)

var (
	toOrderBooks   map[Exchange]chan Order
	fromOrderBooks map[Exchange]chan Fill
	fillBuses      map[Exchange]*FillBus
)

//setupOrderBooks populates the toOrderBooks global map
func setupOrderBooks() {

	toOrderBooks = make(map[Exchange]chan Order)
	fromOrderBooks = make(map[Exchange]chan Fill)

	btcUsd := NewBookManager(BTCUSD)
	btcUsdIn := make(chan Order)
	btcUsdOut := make(chan Fill)

	toOrderBooks[btcUsd.Exchange] = btcUsdIn
	fromOrderBooks[btcUsd.Exchange] = btcUsdOut

	go btcUsd.Run(btcUsdIn, btcUsdOut)

	btcLtc := NewBookManager(BTCLTC)
	btcLtcIn := make(chan Order)
	btcLtcOut := make(chan Fill)

	toOrderBooks[btcLtc.Exchange] = btcLtcIn
	fromOrderBooks[btcLtc.Exchange] = btcLtcOut

	go btcLtc.Run(btcLtcIn, btcLtcOut)

	btcDoge := NewBookManager(BTCDOGE)
	btcDogeIn := make(chan Order)
	btcDogeOut := make(chan Fill)

	toOrderBooks[btcDoge.Exchange] = btcDogeIn
	fromOrderBooks[btcDoge.Exchange] = btcDogeOut

	go btcDoge.Run(btcDogeIn, btcDogeOut)

	btcXmr := NewBookManager(BTCXMR)
	btcXmrIn := make(chan Order)
	btcXmrOut := make(chan Fill)

	toOrderBooks[btcXmr.Exchange] = btcXmrIn
	fromOrderBooks[btcXmr.Exchange] = btcXmrOut

	go btcXmr.Run(btcXmrIn, btcXmrOut)
	fmt.Println("Finished starting channels")

}

func setupOrderBuses() {

	btcUsdBus := NewFillBus()
	go btcUsdBus.Run(fromOrderBooks[BTCUSD])

	btcLtcBus := NewFillBus()
	go btcLtcBus.Run(fromOrderBooks[BTCLTC])

	btcDogeBus := NewFillBus()
	go btcDogeBus.Run(fromOrderBooks[BTCDOGE])

	btcXmrBus := NewFillBus()
	go btcXmrBus.Run(fromOrderBooks[BTCXMR])

	fillBuses = make(map[Exchange]*FillBus)
	fillBuses[BTCUSD] = btcUsdBus
	fillBuses[BTCLTC] = btcLtcBus
	fillBuses[BTCDOGE] = btcDogeBus
	fillBuses[BTCXMR] = btcXmrBus
}

func main() {

	setupOrderBooks()
	setupOrderBuses()

	resolver := NewAccountResolver()
	go resolver.Run(fillBuses)

	go func() { log.Println(http.ListenAndServe("localhost:6060", nil)) }()

	r := mux.NewRouter()

	r.HandleFunc("/health", HealthCheckHandler)
	r.HandleFunc("/order", orderHandler).Methods("POST")
	r.HandleFunc("/cancel", cancelHandler).Methods("POST")
	r.HandleFunc("/order-status/{userId}/{exchange}", orderStatusHandler).Methods("GET")
	r.HandleFunc("/stream/fills/{exchange}", FillsWebsocketHandler).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r))
}

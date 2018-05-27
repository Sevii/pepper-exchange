package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

type Account struct {
	UserId         string
	balanceUSD     int // usd
	balanceSatoshi int // btc
	balanceXMR     int
	balanceDOGE    int
	orders         []Order
	fills          []Fill
}

var accounts map[string]Account

//AccountResolver watches fill updates from the exchange and keeps user accounts up to date
// Writes up-to-date accounts to redis
type AccountResolver struct {
	redisClient *redis.Client
	fillStreams map[Exchange]chan Fill
}

func NewAccountResolver() *AccountResolver {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>

	chans := make(map[Exchange]chan Fill)
	chans[BTCUSD] = make(chan Fill)
	chans[BTCLTC] = make(chan Fill)
	chans[BTCDOGE] = make(chan Fill)
	chans[BTCXMR] = make(chan Fill)

	return &AccountResolver{redisClient: client, fillStreams: chans}
}

func resolveFill(fill Fill) {
	//todo
}

func (r *AccountResolver) Run(buses map[Exchange]*FillBus) {
	buses[BTCUSD].subscribe(r.fillStreams[BTCUSD])
	buses[BTCLTC].subscribe(r.fillStreams[BTCLTC])
	buses[BTCDOGE].subscribe(r.fillStreams[BTCDOGE])
	buses[BTCXMR].subscribe(r.fillStreams[BTCXMR])

	for { //infinite loop
		select {
		case usd := <-r.fillStreams[BTCUSD]:
			fmt.Println("Fill from: ", usd.Participants[0].UserId, usd.Participants[1].UserId)
			// fmt.Println(usd)
			// fmt.Println("")
		case ltc := <-r.fillStreams[BTCLTC]:
			fmt.Println("Fill from: ", ltc.Participants[0].UserId, ltc.Participants[1].UserId)
			// fmt.Println(ltc)
			// fmt.Println("")
		case doge := <-r.fillStreams[BTCDOGE]:
			fmt.Println("Fill from: ", doge.Participants[0].UserId, doge.Participants[1].UserId)
			// fmt.Println(doge)
			// fmt.Println("")
		case xmr := <-r.fillStreams[BTCXMR]:
			fmt.Println("Fill from: ", xmr.Participants[0].UserId, xmr.Participants[1].UserId)
			// fmt.Println(xmr)
			// fmt.Println("")
		}
	}
}

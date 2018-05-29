package main

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
)

var latestFills []Fill

const (
	usdLatest  = "usd.market.price"
	ltcLatest  = "ltc.market.price"
	dogeLatest = "doge.market.price"
	xmrLatest  = "xmr.market.price"
)

type MarketData struct {
	USDPrice  int    `json:"usdPrice"`
	LTCPrice  int    `json:"ltcPrice"`
	DOGEPrice int    `json:"dogePrice"`
	XMRPrice  int    `json:"xmrPrice"`
	LastFills []Fill `json:"lastFills"`
}

func RunMarketData(buses map[Exchange]*FillBus) {
	setupRedisMarketData()
	chans := make(map[Exchange]chan Fill)
	chans[BTCUSD] = make(chan Fill)
	chans[BTCLTC] = make(chan Fill)
	chans[BTCDOGE] = make(chan Fill)
	chans[BTCXMR] = make(chan Fill)

	buses[BTCUSD].subscribe(chans[BTCUSD])
	buses[BTCLTC].subscribe(chans[BTCLTC])
	buses[BTCDOGE].subscribe(chans[BTCDOGE])
	buses[BTCXMR].subscribe(chans[BTCXMR])

	for { //infinite loop
		select {
		case usd := <-chans[BTCUSD]:
			setLatestPrice(usdLatest, usd.Price)
			updateLatestFills(usd)
		case ltc := <-chans[BTCLTC]:
			setLatestPrice(ltcLatest, ltc.Price)
			updateLatestFills(ltc)
		case doge := <-chans[BTCDOGE]:
			setLatestPrice(dogeLatest, doge.Price)
			updateLatestFills(doge)
		case xmr := <-chans[BTCXMR]:
			setLatestPrice(xmrLatest, xmr.Price)
			updateLatestFills(xmr)
		}
	}
}

func getMarketData() (MarketData, error) {
	fillsJson, err := redisClient.Get("latestFills").Result()
	if err != nil {
		return MarketData{}, err
	}

	// Deserialize the order
	var recentFills []Fill
	err = json.Unmarshal([]byte(fillsJson), &recentFills)
	if err != nil {
		recentFills = make([]Fill, 1)
		log.Println("Failed to get recent fills from redis")
	}

	usdMarketString, err := redisClient.Get("usd.market.price").Result()
	if err != nil {
		return MarketData{}, err
	}
	usdMarket, _ := strconv.Atoi(usdMarketString)

	ltcMarketString, err := redisClient.Get("ltc.market.price").Result()
	if err != nil {
		return MarketData{}, err
	}
	ltcMarket, _ := strconv.Atoi(ltcMarketString)

	dogeMarketString, err := redisClient.Get("doge.market.price").Result()
	if err != nil {
		return MarketData{}, err
	}
	dogeMarket, _ := strconv.Atoi(dogeMarketString)

	xmrMarketString, err := redisClient.Get("xmr.market.price").Result()
	if err != nil {
		return MarketData{}, err
	}
	xmrMarket, _ := strconv.Atoi(xmrMarketString)

	return MarketData{
		USDPrice:  usdMarket,
		LTCPrice:  ltcMarket,
		DOGEPrice: dogeMarket,
		XMRPrice:  xmrMarket,
		LastFills: recentFills}, nil

}

func updateLatestFills(fill Fill) error {
	if len(latestFills) < 21 {
		latestFills = append(latestFills, fill)
	} else {
		jsonFills, err := json.Marshal(latestFills)
		if err != nil {
			return errors.New("Cannot convert to JSON")
		}
		err = redisClient.Set("latestFills", jsonFills, 0).Err()
		if err != nil {
			return err
		}

		latestFills = []Fill{}
	}
	return nil
}

func setLatestPrice(coinPath string, newBalance int) error {

	err := redisClient.Set(coinPath, newBalance, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func setupRedisMarketData() {
	setLatestPrice(usdLatest, 0)
	setLatestPrice(ltcLatest, 0)
	setLatestPrice(dogeLatest, 0)
	setLatestPrice(xmrLatest, 0)
	err := redisClient.Set("latestFills", "", 0).Err()
	if err != nil {
		log.Println("Failed to null latestFills on startup")
	}

}

package main

import (
	"fmt"
	"strconv"
)

const (
	pathUSD  = ".usd"
	pathBTC  = ".btc"
	pathLTC  = ".ltc"
	pathXMR  = ".xmr"
	pathDOGE = ".doge"
)

type Account struct {
	UserId string
	USD    int
	LTC    int
	BTC    int
	XMR    int
	DOGE   int
	// orders      []Order
	// fills       []Fill
}

//AccountResolver watches fill updates from the exchange and keeps user accounts up to date
// Writes up-to-date accounts to redis
type AccountResolver struct {
	fillStreams map[Exchange]chan Fill
	accounts    map[string]Account
}

func NewAccountResolver() *AccountResolver {

	chans := make(map[Exchange]chan Fill)
	chans[BTCUSD] = make(chan Fill)
	chans[BTCLTC] = make(chan Fill)
	chans[BTCDOGE] = make(chan Fill)
	chans[BTCXMR] = make(chan Fill)

	accountMap := make(map[string]Account)

	return &AccountResolver{fillStreams: chans, accounts: accountMap}
}

//setupAccounts sets up the account map
func (a *AccountResolver) setupAccounts() {
	users := []string{"BOB", "ALICE", "ROBODOG", "KID1", "KID2", "KID3", "KID4", "OTHERKID"}
	for _, name := range users {
		a.accounts[name] = Account{
			UserId: name,
			USD:    10000,
			BTC:    0,
			DOGE:   0,
			XMR:    0,
			// orders:      make([]Order, 3),
			// fills:       make([]Fill, 3),
		}
	}

}

func (a AccountResolver) initiate() {
	a.setupAccounts()
	a.setupRedis()
}

func (a AccountResolver) setupRedis() error {
	for _, user := range a.accounts {
		usd := user.UserId + ".usd"
		err := redisClient.Set(usd, 10000, 0).Err()
		if err != nil {
			return err
		}

		btc := user.UserId + pathBTC
		ltc := user.UserId + pathLTC
		xmr := user.UserId + pathXMR
		doge := user.UserId + pathDOGE

		initialAccounts := []string{btc, ltc, xmr, doge}
		for _, coin := range initialAccounts {
			err := redisClient.Set(coin, 0, 0).Err()
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func resolveFill(fill Fill) {
	switch fill.Exchange {
	case BTCUSD:
		handleTrade(fill, pathUSD)
	case BTCLTC:
		handleTrade(fill, pathLTC)
	case BTCDOGE:
		handleTrade(fill, pathDOGE)
	case BTCXMR:
		handleTrade(fill, pathXMR)
	default:
		//Do nothing
	}
}

func handleTrade(fill Fill, coinPath string) {
	for _, participant := range fill.Participants {
		if participant.Direction == ASK {
			//Sold number USD
			soldCoin := fill.Number
			// at price * number BTC
			btcRecieved := fill.Price * fill.Number

			oldCoinBalance, err := getBalance(participant.UserId, coinPath)
			if err != nil {
				fmt.Println("Failed to update account status!!")
				return
			}
			oldBtcBalance, err := getBalance(participant.UserId, pathBTC)
			if err != nil {
				fmt.Println("Failed to update account status!!")
				return
			}

			//Update balances
			err = setBalance(participant.UserId, pathUSD, oldCoinBalance-soldCoin)
			if err != nil {
				fmt.Println("Failed to update account status!!")
				return
			}
			err = setBalance(participant.UserId, pathBTC, oldBtcBalance+btcRecieved)
			if err != nil {
				fmt.Println("Failed to update account status!!")
				return
			}

		} else if participant.Direction == BID {
			//Bought number USD
			boughtCoin := fill.Number
			// at price * number BTC
			btcPaid := fill.Price * fill.Number

			oldCoinBalance, err := getBalance(participant.UserId, coinPath)
			if err != nil {
				fmt.Println("Failed to update account status!!")
				return
			}
			oldBtcBalance, err := getBalance(participant.UserId, pathBTC)
			if err != nil {
				fmt.Println("Failed to update account status!!")
				return
			}

			err = setBalance(participant.UserId, pathUSD, oldCoinBalance+boughtCoin)
			if err != nil {
				fmt.Println("Failed to update account status!!")
				return
			}
			err = setBalance(participant.UserId, pathBTC, oldBtcBalance-btcPaid)
			if err != nil {
				fmt.Println("Failed to update account status!!")
				return
			}
		}
	}
}

func setBalance(userId string, coinPath string, newBalance int) error {

	err := redisClient.Set(userId+coinPath, newBalance, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func getBalance(userId string, coinPath string) (int, error) {
	balanceString, err := redisClient.Get(userId + coinPath).Result()
	if err != nil {
		return 0, err
	}
	balance, _ := strconv.Atoi(balanceString)
	return balance, nil
}

func (r *AccountResolver) Run(buses map[Exchange]*FillBus) {
	buses[BTCUSD].subscribe(r.fillStreams[BTCUSD])
	buses[BTCLTC].subscribe(r.fillStreams[BTCLTC])
	buses[BTCDOGE].subscribe(r.fillStreams[BTCDOGE])
	buses[BTCXMR].subscribe(r.fillStreams[BTCXMR])

	for { //infinite loop
		select {
		case usd := <-r.fillStreams[BTCUSD]:
			resolveFill(usd)
		case ltc := <-r.fillStreams[BTCLTC]:
			resolveFill(ltc)
		case doge := <-r.fillStreams[BTCDOGE]:
			resolveFill(doge)
		case xmr := <-r.fillStreams[BTCXMR]:
			resolveFill(xmr)
		}
	}
}

func getAccountStatusRedis(userId string) (Account, error) {
	usdBalance, err := redisClient.Get(userId + pathUSD).Result()
	if err != nil {
		return Account{}, err
	}

	btcBalance, err := redisClient.Get(userId + pathBTC).Result()
	if err != nil {
		return Account{}, err
	}

	ltcBalance, err := redisClient.Get(userId + pathLTC).Result()
	if err != nil {
		return Account{}, err
	}

	dogeBalance, err := redisClient.Get(userId + pathDOGE).Result()
	if err != nil {
		return Account{}, err
	}

	xmrBalance, err := redisClient.Get(userId + pathXMR).Result()
	if err != nil {
		return Account{}, err
	}

	usd, _ := strconv.Atoi(usdBalance)
	btc, _ := strconv.Atoi(btcBalance)
	ltc, _ := strconv.Atoi(ltcBalance)
	doge, _ := strconv.Atoi(dogeBalance)
	xmr, _ := strconv.Atoi(xmrBalance)

	return Account{UserId: userId,
		USD:  int(usd),
		BTC:  int(btc),
		LTC:  int(ltc),
		DOGE: int(doge),
		XMR:  int(xmr),
	}, nil
}

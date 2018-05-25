package main

import (
	"fmt"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"log"
)

// Manager maintains the Orderbook
type BookManager struct {
	logger   log.Logger
	Exchange Exchange
	book     *rbt.Tree
}

// NewManager is the Book Manager constructor
func NewBookManager(exchange Exchange) BookManager {
	return BookManager{
		book:     rbt.NewWithIntComparator(),
		Exchange: exchange,
	}
}

func (d *BookManager) Run(in <-chan Order, out chan<- Fill) {

	for order := range in {
		switch order.Direction {
		case ASK:
			//Ask things
			fmt.Println("ASK")
		case BID:
			//Bid Operations
			fmt.Println("BID")
		case CANCEL:
			//Cancel an order
			fmt.Println("CANCEL")
		default:
			//Drop the message
			fmt.Println("Invalid Order Type")
		}
	}
}

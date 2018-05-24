package main

import (
	"fmt"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"log"
)

// NewManager is the Book Manager constructor
func NewBookManager() *BookManager {
	return &BookManager{
		book: rbt.NewWithIntComparator(),
	}
}

func (d *BookManager) Run(in <-chan Order, out chan<- Fill) {

	for order := range in {
		switch order.Direction {
		case ASK:
			//Ask things
			fmt.Println("ASK")
		case BID:
			fmt.Println("BID")
		case CANCEL:
			fmt.Println("CANCEL")
		default:
			//Drop the message
			fmt.Println("Invalid Order Type")
		}
	}
}

// Manager maintains the Orderbook
type BookManager struct {
	logger log.Logger

	in   <-chan Order
	out  chan<- Fill
	book *rbt.Tree
}

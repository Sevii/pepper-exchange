package main

import (
	"fmt"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
	uuid "github.com/satori/go.uuid"
	"log"
)

type TreeNode struct {
	depth  int //
	orders []Order
}

func NewTreeNode() TreeNode {
	return TreeNode{
		depth:  0,
		orders: make([]Order, 0),
	}
}

func (n TreeNode) add(ord Order) TreeNode {
	n.orders = append(n.orders, ord)
	return n
}

func (n TreeNode) String() string {
	return fmt.Sprint(n.orders)
}

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
			fmt.Println("ASK CHAN")
			addOrder(d.book, order)

		case BID:
			//Bid Operations
			fmt.Println("BID CHAN")
			addOrder(d.book, order)
		case CANCEL:
			//Cancel an order
			fmt.Println("CANCEL CHAN")
			addOrder(d.book, order)
		default:
			//Drop the message
			fmt.Println("Invalid Order Type")
		}
	}
}

//executeOrder walk the orderbook and match asks and bids that can fill
func executeOrder(book *rbt.Tree, ord Order) (Order, error) {
	return Order{}, nil
}

//addOrder adds an order that cannot be filled any further to the orderbook
func addOrder(book *rbt.Tree, ord Order) (Order, error) {
	treeNode, ok := book.Get(ord.Price)
	if !ok {
		node := NewTreeNode()
		book.Put(ord.Price, node.add(ord))
		fmt.Println("Added array")
	} else {
		book.Put(ord.Price, treeNode.(TreeNode).add(ord))
		fmt.Println("Added order")
	}

	fmt.Println(book)
	iter := book.Iterator()
	for iter.Next() == true {
		fmt.Println(iter.Value())
		fmt.Println()
	}

	return Order{}, nil
}

//cancelOrder cancels (deletes) an outstanding order from the orderbook
func cancelOrder(book *rbt.Tree, orderId uuid.UUID) bool {
	return false
}

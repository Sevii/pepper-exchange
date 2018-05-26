package main

import (
	"fmt"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
	uuid "github.com/satori/go.uuid"
	"log"
	"sort"
)

type TreeNode struct {
	depth  int //
	orders map[uuid.UUID]Order
}

func NewTreeNode() *TreeNode {
	return &TreeNode{
		depth:  0,
		orders: make(map[uuid.UUID]Order),
	}
}

func (n *TreeNode) upsert(ord Order) {
	n.orders[ord.ID] = ord
}

func (n *TreeNode) delete(id uuid.UUID) {
	delete(n.orders, id)
}

func (n *TreeNode) get(id uuid.UUID) Order {
	return n.orders[id]
}

func (n TreeNode) sortedOrders() []Order {
	orders := make([]Order, 0)
	for _, v := range n.orders {
		orders = append(orders, v)
	}

	sort.Slice(orders[:], func(i, j int) bool {
		return orders[i].Timestamp < orders[j].Timestamp
	})

	return orders
}

func (n TreeNode) String() string {
	vals := make([]Order, len(n.orders))
	for _, v := range n.orders {
		vals = append(vals, v)
	}
	return fmt.Sprint(vals)
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
func executeOrder(book *rbt.Tree, ord Order) (Order, []Fill) {
	var fills []Fill

	if ord.Direction == BID {
		// start left
		it := book.Iterator()

		for it.Begin(); it.Next(); {
			nodePrice, node := it.Key().(int), it.Value().(*TreeNode)

			//Check price
			if ord.Price >= nodePrice {
				// Have to append to top level variable but am dealing with scoped binding as well :=
				nodeOrderResult, nodeFills := matchNode(node, ord)
				ord = nodeOrderResult

				fills = append(fills, nodeFills...)
			} else {
				//skip this node, too expensive (The cheapest ask could be higher than this bid)
				continue
			}

			if ord.NumberOutstanding == 0 {
				// if we have 0 outstanding we can quit
				break
			}

			fmt.Println("Node Price: \n", nodePrice)
			fmt.Println("Node: \n", node)

		}
		return ord, fills
	} else if ord.Direction == ASK {
		// start left
		it := book.Iterator()

		for it.End(); it.Prev(); {
			nodePrice, node := it.Key().(int), it.Value().(*TreeNode)

			//Check price
			if ord.Price <= nodePrice {
				// Have to append to top level variable but am dealing with scoped binding as well :=
				nodeOrderResult, nodeFills := matchNode(node, ord)
				ord = nodeOrderResult

				fills = append(fills, nodeFills...)
			} else {
				//skip this node, too expensive (The cheapest ask could be higher than this bid)
				continue
			}

			if ord.NumberOutstanding == 0 {
				// if we have 0 outstanding we can quit
				break
			}

			fmt.Println("Node Price: \n", nodePrice)
			fmt.Println("Node: \n", node)

		}
		return ord, fills
	} else {
		// Not a valid bid/ask
	}

	return Order{}, nil
}

//matchNode takes an order and fills it against a node, NOT IDEMPOTENT
func matchNode(node *TreeNode, ord Order) (Order, []Fill) {
	//We only deal with ask and bid
	if ord.Direction == CANCEL || ord.Direction == INVALID_OPERATION {
		return ord, []Fill{}
	}

	orders := node.sortedOrders()
	activeOrder := ord
	var fills []Fill

	for _, oldOrder := range orders {
		if activeOrder.Direction != oldOrder.Direction {
			fmt.Println("ActiveOrd Outstanding: ", activeOrder.NumberOutstanding)
			fmt.Println("Old order outstanding: ", oldOrder.Number)

			// If the current order can fill new order
			if ord.NumberOutstanding <= oldOrder.NumberOutstanding {
				part := []Order{activeOrder, oldOrder}
				closed := []Order{activeOrder}
				if oldOrder.NumberOutstanding-ord.NumberOutstanding == 0 {
					closed = append(closed, oldOrder)
					node.delete(oldOrder.ID)
				} else { // Update old order
					oldRemaining := oldOrder.NumberOutstanding - activeOrder.NumberOutstanding
					oldOrder.NumberOutstanding = oldRemaining
					node.upsert(oldOrder)
				}

				fill := NewFill(activeOrder.Exchange, activeOrder.NumberOutstanding, oldOrder.Price, part, closed)

				//Order is filled
				activeOrder.NumberOutstanding = 0
				fills = append(fills, fill)

			} else { // If the current order is to small to fill the new order
				//How do we delete the old order?
				node.delete(oldOrder.ID)

				part := []Order{activeOrder, oldOrder}
				closed := []Order{oldOrder}
				fill := NewFill(activeOrder.Exchange, oldOrder.NumberOutstanding, oldOrder.Price, part, closed)

				activeOrder.NumberOutstanding = activeOrder.NumberOutstanding - oldOrder.NumberOutstanding
				fills = append(fills, fill)

			}

		}
	}

	return activeOrder, fills
}

//addOrder adds an order that cannot be filled any further to the orderbook
func addOrder(book *rbt.Tree, ord Order) (Order, error) {
	treeNode, ok := book.Get(ord.Price)
	if !ok {
		node := NewTreeNode()
		node.upsert(ord)

		book.Put(ord.Price, node)

	} else {
		node := treeNode.(*TreeNode)
		node.upsert(ord)
		book.Put(ord.Price, treeNode)
	}
	fmt.Println(book)

	return Order{}, nil
}

//cancelOrder cancels (deletes) an outstanding order from the orderbook
func cancelOrder(book *rbt.Tree, orderId uuid.UUID) Fill {
	var ord Order
	it := book.Iterator()
	for it.Begin(); it.Next(); {
		_, node := it.Key().(int), it.Value().(*TreeNode)
		ord = node.get(orderId)
		node.delete(orderId)

	}
	part := []Order{ord}
	closed := []Order{ord}
	fill := NewFill(ord.Exchange, ord.NumberOutstanding, ord.Price, part, closed)
	return fill
}

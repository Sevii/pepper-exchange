package main

import (
	rbt "github.com/emirpasic/gods/trees/redblacktree"
	uuid "github.com/satori/go.uuid"
	"testing"
	"time"
)

func TestMatchNode(t *testing.T) {
	testCases := []struct {
		title         string
		nodeOrders    []Order
		activeOrder   Order
		expectedFills int
		expectedDepth int
	}{
		{
			title:         "Test Match an order against a deep node",
			expectedFills: 3,
			expectedDepth: 0,
			activeOrder: Order{
				ID:                uuid.NewV4(),
				Direction:         BID,
				Exchange:          BTCDOGE,
				Number:            1000,
				NumberOutstanding: 1000,
				Price:             3000,
				Timestamp:         time.Now().Nanosecond(),
			},
			nodeOrders: []Order{
				Order{
					ID:                uuid.NewV4(),
					Direction:         ASK,
					Exchange:          BTCDOGE,
					Number:            1000,
					NumberOutstanding: 1000,
					Price:             3000,
					Timestamp:         time.Now().Nanosecond(),
				},

				Order{
					ID:                uuid.NewV4(),
					Direction:         ASK,
					Exchange:          BTCDOGE,
					Number:            1000,
					NumberOutstanding: 1000,
					Price:             3000,
					Timestamp:         time.Now().Nanosecond(),
				},

				Order{
					ID:                uuid.NewV4(),
					Direction:         ASK,
					Exchange:          BTCDOGE,
					Number:            1000,
					NumberOutstanding: 1000,
					Price:             3000,
					Timestamp:         time.Now().Nanosecond(),
				},
			},
		},
		{
			title:         "Test Match an order against a shallow node",
			expectedFills: 3,
			expectedDepth: 1000,
			activeOrder: Order{
				ID:                uuid.NewV4(),
				Direction:         BID,
				Exchange:          BTCDOGE,
				Number:            4000,
				NumberOutstanding: 4000,
				Price:             3000,
				Timestamp:         time.Now().Nanosecond(),
			},
			nodeOrders: []Order{
				Order{
					ID:                uuid.NewV4(),
					Direction:         ASK,
					Exchange:          BTCDOGE,
					Number:            1000,
					NumberOutstanding: 1000,
					Price:             3000,
					Timestamp:         time.Now().Nanosecond(),
				},

				Order{
					ID:                uuid.NewV4(),
					Direction:         ASK,
					Exchange:          BTCDOGE,
					NumberOutstanding: 1000,
					Price:             3000,
					Timestamp:         time.Now().Nanosecond(),
				},

				Order{
					ID:                uuid.NewV4(),
					Direction:         ASK,
					Exchange:          BTCDOGE,
					Number:            1000,
					NumberOutstanding: 1000,
					Price:             3000,
					Timestamp:         time.Now().Nanosecond(),
				},
			},
		},
		{
			title:         "Test Match an order against a short node",
			expectedFills: 1,
			expectedDepth: 3000,
			activeOrder: Order{
				ID:                uuid.NewV4(),
				Direction:         BID,
				Exchange:          BTCDOGE,
				Number:            4000,
				NumberOutstanding: 4000,
				Price:             3000,
				Timestamp:         time.Now().Nanosecond(),
			},
			nodeOrders: []Order{
				Order{
					ID:                uuid.NewV4(),
					Direction:         ASK,
					Exchange:          BTCDOGE,
					Number:            1000,
					NumberOutstanding: 1000,
					Price:             3000,
					Timestamp:         time.Now().Nanosecond(),
				},
			},
		},
		{
			title:         "Test Match an order against a matching direction",
			expectedFills: 0,
			expectedDepth: 4000,
			activeOrder: Order{
				ID:                uuid.NewV4(),
				Direction:         BID,
				Exchange:          BTCDOGE,
				Number:            4000,
				NumberOutstanding: 4000,
				Price:             3000,
				Timestamp:         time.Now().Nanosecond(),
			},
			nodeOrders: []Order{
				Order{
					ID:                uuid.NewV4(),
					Direction:         BID,
					Exchange:          BTCDOGE,
					Number:            1000,
					NumberOutstanding: 1000,
					Price:             3000,
					Timestamp:         time.Now().Nanosecond(),
				},
			},
		},
	}

	for _, testCase := range testCases {

		node := NewTreeNode()
		for _, nodeOrder := range testCase.nodeOrders {
			node.upsert(nodeOrder)
		}

		ord, fills := matchNode(node, testCase.activeOrder)

		actualFills := len(fills)
		if actualFills != testCase.expectedFills || ord.NumberOutstanding != testCase.expectedDepth {

			t.Errorf("%v .\n expected fills: %v, actual fills: %v, expected depth: %v, outstanding depth: %v ",
				testCase.title,
				testCase.expectedFills,
				actualFills,
				testCase.expectedDepth,
				ord.NumberOutstanding)
		}
	}
}

func TestTreeNode(t *testing.T) {
	testCases := []struct {
		title  string
		inputs []Order
		depth  int
	}{
		{
			title: "Inserts into Node",
			depth: 3,
			inputs: []Order{
				Order{
					ID:        uuid.NewV4(),
					Direction: ASK,
					Exchange:  BTCDOGE,
					Number:    1000,
					Price:     3000,
					Timestamp: time.Now().Nanosecond(),
				},

				Order{
					ID:        uuid.NewV4(),
					Direction: ASK,
					Exchange:  BTCDOGE,
					Number:    1000,
					Price:     3000,
					Timestamp: time.Now().Nanosecond(),
				},

				Order{
					ID:        uuid.NewV4(),
					Direction: ASK,
					Exchange:  BTCDOGE,
					Number:    1000,
					Price:     3000,
					Timestamp: time.Now().Nanosecond(),
				},
			},
		},
	}

	for _, testCase := range testCases {

		node := NewTreeNode()
		for _, input := range testCase.inputs {
			node.upsert(input)
		}

		depth := len(node.sortedOrders())
		if depth != testCase.depth {

			t.Errorf("%v. expected length: %v actual length: %v length of orders: %v",
				testCase.title,
				testCase.depth,
				depth,
				len(node.orders))
		}
	}

}

func TestAddOrder(t *testing.T) {

	testCases := []struct {
		title          string
		inputs         []Order
		price          int
		expectedLength int
		depth          int
	}{
		{
			title:          "Test depth after adding matching prices.",
			expectedLength: 3,
			price:          3000,
			depth:          3000,
			inputs: []Order{
				Order{
					ID:        uuid.NewV4(),
					Direction: ASK,
					Exchange:  BTCDOGE,
					Number:    1000,
					Price:     3000,
					Timestamp: time.Now().Nanosecond(),
				},

				Order{
					ID:        uuid.NewV4(),
					Direction: ASK,
					Exchange:  BTCDOGE,
					Number:    1000,
					Price:     3000,
					Timestamp: time.Now().Nanosecond(),
				},

				Order{
					ID:        uuid.NewV4(),
					Direction: ASK,
					Exchange:  BTCDOGE,
					Number:    1000,
					Price:     3000,
					Timestamp: time.Now().Nanosecond(),
				},
			},
		},
		{
			title:          "Test depth after adding different prices.",
			expectedLength: 1,
			price:          1,
			depth:          3000,
			inputs: []Order{
				Order{
					ID:        uuid.NewV4(),
					Direction: ASK,
					Exchange:  BTCDOGE,
					Number:    1000,
					Price:     1,
					Timestamp: time.Now().Nanosecond(),
				},

				Order{
					ID:        uuid.NewV4(),
					Direction: ASK,
					Exchange:  BTCDOGE,
					Number:    1000,
					Price:     2,
					Timestamp: time.Now().Nanosecond(),
				},

				Order{
					ID:        uuid.NewV4(),
					Direction: ASK,
					Exchange:  BTCDOGE,
					Number:    1000,
					Price:     3,
					Timestamp: time.Now().Nanosecond(),
				},
			},
		},
	}

	for _, testCase := range testCases {

		orderbook := rbt.NewWithIntComparator()
		for _, input := range testCase.inputs {
			addOrder(orderbook, input)
		}

		treeNode, _ := orderbook.Get(testCase.price)
		depth := len(treeNode.(*TreeNode).sortedOrders())
		if depth != testCase.expectedLength {

			t.Errorf("%v. expected length: %v actual length: %v",
				testCase.title,
				testCase.expectedLength,
				depth)
		}
	}
}

func TestExecuteOrder(t *testing.T) {
	testCases := []struct {
		title          string
		nodeOrders     []Order
		activeOrder    Order
		expectedFills  int
		expectedDepth  int
		expectedClosed int
	}{
		{
			title:          "Test Execute a Bid order against single node",
			expectedFills:  1,
			expectedDepth:  0,
			expectedClosed: 2,
			activeOrder: Order{
				ID:                uuid.NewV4(),
				Direction:         BID,
				Exchange:          BTCDOGE,
				Number:            1500,
				NumberOutstanding: 1000,
				Price:             4000,
				Timestamp:         time.Now().Nanosecond(),
			},
			nodeOrders: []Order{
				Order{
					ID:                uuid.NewV4(),
					Direction:         ASK,
					Exchange:          BTCDOGE,
					Number:            1000,
					NumberOutstanding: 1000,
					Price:             3000,
					Timestamp:         time.Now().Nanosecond(),
				},

				Order{
					ID:                uuid.NewV4(),
					Direction:         ASK,
					Exchange:          BTCDOGE,
					Number:            1000,
					NumberOutstanding: 1000,
					Price:             4000,
					Timestamp:         time.Now().Nanosecond(),
				},

				Order{
					ID:                uuid.NewV4(),
					Direction:         ASK,
					Exchange:          BTCDOGE,
					Number:            1000,
					NumberOutstanding: 1000,
					Price:             5000,
					Timestamp:         time.Now().Nanosecond(),
				},
			},
		},
		{
			title:          "Test Execute a Bid against multiple nodes",
			expectedFills:  1,
			expectedDepth:  0,
			expectedClosed: 2,
			activeOrder: Order{
				ID:                uuid.NewV4(),
				Direction:         BID,
				Exchange:          BTCDOGE,
				Number:            3500,
				NumberOutstanding: 1000,
				Price:             7777,
				Timestamp:         time.Now().Nanosecond(),
			},
			nodeOrders: []Order{
				Order{
					ID:                uuid.NewV4(),
					Direction:         ASK,
					Exchange:          BTCDOGE,
					Number:            1000,
					NumberOutstanding: 1000,
					Price:             3000,
					Timestamp:         time.Now().Nanosecond(),
				},

				Order{
					ID:                uuid.NewV4(),
					Direction:         ASK,
					Exchange:          BTCDOGE,
					Number:            1000,
					NumberOutstanding: 1000,
					Price:             4000,
					Timestamp:         time.Now().Nanosecond(),
				},

				Order{
					ID:                uuid.NewV4(),
					Direction:         ASK,
					Exchange:          BTCDOGE,
					Number:            1000,
					NumberOutstanding: 1000,
					Price:             5000,
					Timestamp:         time.Now().Nanosecond(),
				},
			},
		},
		{
			title:          "Test Execute a Bid against multiple nodes with outstanding",
			expectedFills:  3,
			expectedDepth:  500,
			expectedClosed: 3,
			activeOrder: Order{
				ID:                uuid.NewV4(),
				Direction:         BID,
				Exchange:          BTCDOGE,
				Number:            2000,
				NumberOutstanding: 2000,
				Price:             7777,
				Timestamp:         time.Now().Nanosecond(),
			},
			nodeOrders: []Order{
				Order{
					ID:                uuid.NewV4(),
					Direction:         ASK,
					Exchange:          BTCDOGE,
					Number:            500,
					NumberOutstanding: 500,
					Price:             3000,
					Timestamp:         time.Now().Nanosecond(),
				},

				Order{
					ID:                uuid.NewV4(),
					Direction:         ASK,
					Exchange:          BTCDOGE,
					Number:            500,
					NumberOutstanding: 500,
					Price:             4000,
					Timestamp:         time.Now().Nanosecond(),
				},

				Order{
					ID:                uuid.NewV4(),
					Direction:         ASK,
					Exchange:          BTCDOGE,
					Number:            1000,
					NumberOutstanding: 500,
					Price:             5000,
					Timestamp:         time.Now().Nanosecond(),
				},
			},
		},
		{
			title:          "Test Execute a ASK against single node",
			expectedFills:  1,
			expectedDepth:  0,
			expectedClosed: 2,
			activeOrder: Order{
				ID:                uuid.NewV4(),
				Direction:         ASK,
				Exchange:          BTCDOGE,
				Number:            1500,
				NumberOutstanding: 1000,
				Price:             2500,
				Timestamp:         time.Now().Nanosecond(),
			},
			nodeOrders: []Order{
				Order{
					ID:                uuid.NewV4(),
					Direction:         BID,
					Exchange:          BTCDOGE,
					Number:            1000,
					NumberOutstanding: 1000,
					Price:             3000,
					Timestamp:         time.Now().Nanosecond(),
				},

				Order{
					ID:                uuid.NewV4(),
					Direction:         BID,
					Exchange:          BTCDOGE,
					Number:            1000,
					NumberOutstanding: 1000,
					Price:             4000,
					Timestamp:         time.Now().Nanosecond(),
				},

				Order{
					ID:                uuid.NewV4(),
					Direction:         BID,
					Exchange:          BTCDOGE,
					Number:            1000,
					NumberOutstanding: 1000,
					Price:             5000,
					Timestamp:         time.Now().Nanosecond(),
				},
			},
		},
		{
			title:          "Test Execute a ASK against multiple nodes",
			expectedFills:  1,
			expectedDepth:  0,
			expectedClosed: 2,
			activeOrder: Order{
				ID:                uuid.NewV4(),
				Direction:         ASK,
				Exchange:          BTCDOGE,
				Number:            3500,
				NumberOutstanding: 1000,
				Price:             777,
				Timestamp:         time.Now().Nanosecond(),
			},
			nodeOrders: []Order{
				Order{
					ID:                uuid.NewV4(),
					Direction:         BID,
					Exchange:          BTCDOGE,
					Number:            1000,
					NumberOutstanding: 1000,
					Price:             3000,
					Timestamp:         time.Now().Nanosecond(),
				},

				Order{
					ID:                uuid.NewV4(),
					Direction:         BID,
					Exchange:          BTCDOGE,
					Number:            1000,
					NumberOutstanding: 1000,
					Price:             4000,
					Timestamp:         time.Now().Nanosecond(),
				},

				Order{
					ID:                uuid.NewV4(),
					Direction:         BID,
					Exchange:          BTCDOGE,
					Number:            1000,
					NumberOutstanding: 1000,
					Price:             5000,
					Timestamp:         time.Now().Nanosecond(),
				},
			},
		},
		{
			title:          "Test Execute a ASK against multiple nodes with outstanding",
			expectedFills:  3,
			expectedDepth:  500,
			expectedClosed: 3,
			activeOrder: Order{
				ID:                uuid.NewV4(),
				Direction:         ASK,
				Exchange:          BTCDOGE,
				Number:            2000,
				NumberOutstanding: 2000,
				Price:             777,
				Timestamp:         time.Now().Nanosecond(),
			},
			nodeOrders: []Order{
				Order{
					ID:                uuid.NewV4(),
					Direction:         BID,
					Exchange:          BTCDOGE,
					Number:            500,
					NumberOutstanding: 500,
					Price:             3000,
					Timestamp:         time.Now().Nanosecond(),
				},

				Order{
					ID:                uuid.NewV4(),
					Direction:         BID,
					Exchange:          BTCDOGE,
					Number:            500,
					NumberOutstanding: 500,
					Price:             4000,
					Timestamp:         time.Now().Nanosecond(),
				},

				Order{
					ID:                uuid.NewV4(),
					Direction:         BID,
					Exchange:          BTCDOGE,
					Number:            1000,
					NumberOutstanding: 500,
					Price:             5000,
					Timestamp:         time.Now().Nanosecond(),
				},
			},
		},
	}

	for _, testCase := range testCases {

		mgmt := NewBookManager(BTCDOGE)
		book := mgmt.book
		for _, nodeOrder := range testCase.nodeOrders {
			addOrder(book, nodeOrder)
		}

		ord, fills := executeOrder(book, testCase.activeOrder)

		//Fills should only be one
		var closed []Order
		for _, fill := range fills {
			closed = append(closed, fill.Closed...)
		}

		actualClosed := len(closed)

		actualFills := len(fills)
		if actualFills != testCase.expectedFills || ord.NumberOutstanding != testCase.expectedDepth || actualClosed != testCase.expectedClosed {

			t.Errorf("%v .\n expected fills: %v,  actual fills: %v,\n expected depth: %v, outstanding depth: %v,\n expected closed: %v count closed: %v, \n fills %+v",
				testCase.title,
				testCase.expectedFills,
				actualFills,
				testCase.expectedDepth,
				ord.NumberOutstanding,
				testCase.expectedClosed,
				actualClosed,
				fills)
		}
	}
}

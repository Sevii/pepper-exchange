package main

import (
	rbt "github.com/emirpasic/gods/trees/redblacktree"
	uuid "github.com/satori/go.uuid"
	"testing"
	"time"
)

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
	}

	for _, testCase := range testCases {

		orderbook := rbt.NewWithIntComparator()
		for _, input := range testCase.inputs {
			addOrder(orderbook, input)
		}

		treeNode, _ := orderbook.Get(testCase.price)
		depth := len(treeNode.(TreeNode).orders)
		if depth != testCase.expectedLength {

			t.Errorf("%v. expected length: %v actual length: %v",
				testCase.title,
				testCase.expectedLength,
				depth)
		}
	}
}

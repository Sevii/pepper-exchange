package main

import (
	"testing"
)

func TestOperationToString(t *testing.T) {

	testCases := []struct {
		title    string
		input    Operation
		expected string
	}{
		{
			title:    "BID converts to bid",
			input:    BID,
			expected: "BID",
		},
		{
			title:    "ASK converts to ask",
			input:    ASK,
			expected: "ASK",
		},
		{
			title:    "CANCEL converts to cancel",
			input:    CANCEL,
			expected: "CANCEL",
		},
	}

	for _, testCase := range testCases {
		actual := testCase.input.String()
		if actual != testCase.expected {

			t.Errorf("%v. input: %v \nexpected: %v, \nactual: %v",
				testCase.title,
				testCase.input,
				testCase.expected,
				actual)
		}
	}
}

func TestOperationFromStr(t *testing.T) {

	testCases := []struct {
		title    string
		input    string
		expected Operation
	}{
		{
			title:    "bid converts to BID",
			input:    "bid",
			expected: BID,
		},
		{
			title:    "ask converts to ASK",
			input:    "ask",
			expected: ASK,
		},
		{
			title:    "cancel converts to CANCEL",
			input:    "cancel",
			expected: CANCEL,
		},
	}

	for _, testCase := range testCases {
		if OperationFromStr(testCase.input) != testCase.expected {

			t.Errorf("%v. input: %v \nexpected: %v",
				testCase.title,
				testCase.input,
				testCase.expected)
		}
	}

}

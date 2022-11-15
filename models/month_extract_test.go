package models

import (
	"math"
	"testing"
	"time"
)

func roundComparison6Places(value float64) float64 {
	return (math.Ceil(value*1e7) / 1e7) - 1e-7
}

type comparisonTest struct {
	Day                int
	TransactionsInput  []Transaction
	TransactionsOutput []Transaction
}

func TestGetMonthExtract(t *testing.T) {
	// Get a random date
	date := time.Date(2022, 9, 21, 0, 0, 0, 0, time.UTC)

	// Getting the MonthExtract object
	var me MonthExtract
	me.GetMonthExtract(date)

	// Comparing Total values
	expectedInput := 3356.96
	expectedOutput := 23.32
	roundedGotTotalInputs := roundComparison6Places(me.TotalInputs)
	if roundedGotTotalInputs != expectedInput {
		t.Errorf("Invalid TotalInputs, got %v expected %v", roundedGotTotalInputs, expectedInput)
	}
	if me.TotalOutputs != expectedOutput {
		t.Errorf("Invalid TotalOutputs, got %v expected %v", me.TotalOutputs, expectedOutput)
	}

	// Comparing the days gotten
	expectedLenght := 30
	if lenght := len(me.ExtractsPerMonth); lenght != expectedLenght {
		t.Errorf("Invalid ExtractsPerMonth lenght, expected %v got %v", expectedLenght, lenght)
	}

	// Comparing the days gotten
	expectedDays := []comparisonTest{
		{
			Day: 21,
			TransactionsInput: []Transaction{
				{
					Mode:      Pix,
					Receiving: true,
					Recipient: "378.432.324-89",
					Value:     3232.32,
				},
				{
					Mode:      Pix,
					Receiving: true,
					Recipient: "323.432.324-89",
					Value:     123.32,
				},
			},
		},
		{
			Day: 22,
			TransactionsOutput: []Transaction{
				{
					Mode:      Pix,
					Receiving: false,
					Recipient: "378.432.324-89",
					Value:     23.32,
				},
			},
		},
		{
			Day: 23,
			TransactionsInput: []Transaction{
				{
					Mode:      Pix,
					Receiving: true,
					Recipient: "378.431.874-89",
					Value:     1.32,
				},
			},
		},
	}

	for _, test := range expectedDays {
		// Checking for transactions inputs
		for i, transaction := range test.TransactionsInput {
			if expTransaction := me.ExtractsPerMonth[test.Day-1].TransactionInput[i]; transaction != expTransaction {
				t.Errorf("Invalid value for TransactionInput, expected %v got %v", transaction, expTransaction)
			}
		}
		// Checking for transactions outputs
		for i, transaction := range test.TransactionsOutput {
			if expTransaction := me.ExtractsPerMonth[test.Day-1].TransactionOutput[i]; transaction != expTransaction {
				t.Errorf("Invalid value for TransactionInput, expected %v got %v", transaction, expTransaction)
			}
		}
	}
}

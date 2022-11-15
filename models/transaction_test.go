package models

import (
	"testing"
	"time"
)

func TestFetchAndTransformDailyTransaction(t *testing.T) {
	expected := []Transaction{
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
		{
			Mode:      Pix,
			Receiving: false,
			Recipient: "378.432.324-89",
			Value:     23.32,
		},
		{
			Mode:      Pix,
			Receiving: true,
			Recipient: "378.431.874-89",
			Value:     1.32,
		},
	}

	date1 := time.Date(2022, 9, 21, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2022, 9, 22, 0, 0, 0, 0, time.UTC)
	date3 := time.Date(2022, 9, 23, 0, 0, 0, 0, time.UTC)

	transactions := append(
		FetchAndTransformDailyTransactions(date1),
		FetchAndTransformDailyTransactions(date2)...,
	)
	transactions = append(transactions, FetchAndTransformDailyTransactions(date3)...)

	for i := range expected {
		if transactions[i].Mode != expected[i].Mode {
			t.Errorf("Output Mode (%q) not equal to expected (%q)", transactions[i].Mode, expected[i].Mode)
		}
		if transactions[i].Receiving != expected[i].Receiving {
			t.Errorf("%v Output Receiving (%v) not equal to expected (%v)", i, transactions[i].Receiving, expected[i].Receiving)
		}
		if transactions[i].Recipient != expected[i].Recipient {
			t.Errorf("Output Recipient (%q) not equal to expected (%q)", transactions[i].Recipient, expected[i].Recipient)
		}
		if transactions[i].Value != expected[i].Value {
			t.Errorf("Output Value (%v) not equal to expected (%v)", transactions[i].Value, expected[i].Value)
		}
	}
}

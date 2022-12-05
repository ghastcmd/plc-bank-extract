package database

import (
	"fmt"
	"testing"
	"time"
)

func testPrint(date time.Time) {
	trans := FetchDailyTransactions(date)
	for _, tran := range trans {
		fmt.Println(tran)
	}
}

func TestFetchDailyTransactions(t *testing.T) {
	date1 := time.Date(2022, 9, 21, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2022, 9, 22, 0, 0, 0, 0, time.UTC)
	date3 := time.Date(2022, 9, 23, 0, 0, 0, 0, time.UTC)

	transactions := append(
		FetchDailyTransactions(date1),
		FetchDailyTransactions(date2)...,
	)
	transactions = append(transactions, FetchDailyTransactions(date3)...)

	expected := []JsonRespRaw{
		{
			Date:      "2022-09-21",
			Mode:      "PIX",
			Receiving: "true",
			Recipient: "378.432.324-89",
			Value:     3232.32,
		},
		{
			Date:      "2022-09-21",
			Mode:      "PIX",
			Receiving: "true",
			Recipient: "323.432.324-89",
			Value:     123.32,
		},
		{
			Date:      "2022-09-22",
			Mode:      "PIX",
			Receiving: "false",
			Recipient: "378.432.324-89",
			Value:     23.32,
		},
		{
			Date:      "2022-09-23",
			Mode:      "PIX",
			Receiving: "true",
			Recipient: "378.431.874-89",
			Value:     1.32,
		},
	}

	if len(expected) != len(transactions) {
		t.Errorf("Lenghts of expectedInput and de.TransactionInput differ, got %v, expected %v", len(transactions), len(expected))
	}

	for i := range expected {
		if transactions[i].Date != expected[i].Date {
			t.Errorf("%v Output Date (%q) not equal to expected (%q)", i, transactions[i].Date, expected[i].Date)
		}
		if transactions[i].Mode != expected[i].Mode {
			t.Errorf("Output Mode (%q) not equal to expected (%q)", transactions[i].Mode, expected[i].Mode)
		}
		if transactions[i].Receiving != expected[i].Receiving {
			t.Errorf("Output Receiving (%q) not equal to expected (%q)", transactions[i].Receiving, expected[i].Receiving)
		}
		if transactions[i].Recipient != expected[i].Recipient {
			t.Errorf("Output Recipient (%q) not equal to expected (%q)", transactions[i].Recipient, expected[i].Recipient)
		}
		if transactions[i].Value != expected[i].Value {
			t.Errorf("Output Value (%v) not equal to expected (%v)", transactions[i].Value, expected[i].Value)
		}
	}
}

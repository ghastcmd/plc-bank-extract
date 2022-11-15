package controllers

import "testing"

func TestFetchDailyTransactions(t *testing.T) {
	transactions := FetchDailyTransactions()
	expected := []JsonRespRaw{
		{
			Date:      "2022-09-21",
			Mode:      "PIX",
			Receiving: "true",
			Recipient: "378.432.324-89",
			Value:     3232.32,
		},
		{
			Date:      "2022-09-22",
			Mode:      "PIX",
			Receiving: "false",
			Recipient: "378.432.324-89",
			Value:     23.32,
		},
		{
			Date:      "2022-09-21",
			Mode:      "PIX",
			Receiving: "true",
			Recipient: "323.432.324-89",
			Value:     123.32,
		},
		{
			Date:      "2022-09-23",
			Mode:      "PIX",
			Receiving: "true",
			Recipient: "378.431.874-89",
			Value:     1.32,
		},
	}

	for i := range expected {
		if transactions[i].Date != expected[i].Date {
			t.Errorf("Output Date (%q) not equal to expected (%q)", transactions[i].Date, expected[i].Date)
		} else if transactions[i].Mode != expected[i].Mode {
			t.Errorf("Output Mode (%q) not equal to expected (%q)", transactions[i].Mode, expected[i].Mode)
		} else if transactions[i].Receiving != expected[i].Receiving {
			t.Errorf("Output Receiving (%q) not equal to expected (%q)", transactions[i].Receiving, expected[i].Receiving)
		} else if transactions[i].Recipient != expected[i].Recipient {
			t.Errorf("Output Recipient (%q) not equal to expected (%q)", transactions[i].Recipient, expected[i].Recipient)
		} else if transactions[i].Value != expected[i].Value {
			t.Errorf("Output Value (%v) not equal to expected (%v)", transactions[i].Value, expected[i].Value)
		}
	}
}

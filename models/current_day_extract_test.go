package models

import (
	"fmt"
	"testing"
	"time"
)

func TestGetDayExtract(t *testing.T) {
	var de DayExtract
	de.GetDayExtract(time.Date(2022, 9, 21, 0, 0, 0, 0, time.UTC))

	expectedInput := []Transaction{
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
	}

	expectedOutput := []Transaction{}

	if len(expectedInput) != len(de.TransactionInput) {
		t.Errorf("Lenghts of expectedInput and de.TransactionInput differ, got: %v and %v", len(expectedInput), len(de.TransactionInput))
	}

	if len(expectedOutput) != len(de.TransactionOutput) {
		t.Errorf("Lenghts of expectedOutput and de.TransactionOutput differ, got: %v and %v", len(expectedOutput), len(de.TransactionOutput))
	}

	for i := range expectedInput {
		if expectedInput[i].Mode != de.TransactionInput[i].Mode {
			t.Errorf("Output Mode (%v) not equal to (%v)", expectedInput[i].Mode, de.TransactionInput[i].Mode)
		} else if expectedInput[i].Receiving != de.TransactionInput[i].Receiving {
			t.Errorf("Output Receiving (%v) not equal to (%v)", expectedInput[i].Receiving, de.TransactionInput[i].Receiving)
		} else if expectedInput[i].Recipient != de.TransactionInput[i].Recipient {
			t.Errorf("Output Recipient (%v) not equal to (%v)", expectedInput[i].Recipient, de.TransactionInput[i].Recipient)
		} else if expectedInput[i].Value != de.TransactionInput[i].Value {
			t.Errorf("Output Value (%v) not equal to (%v)", expectedInput[i].Value, de.TransactionInput[i].Value)
		}
	}

	for i := range expectedOutput {
		if expectedOutput[i].Mode != de.TransactionOutput[i].Mode {
			t.Errorf("Output Mode (%v) not equal to (%v)", expectedOutput[i].Mode, de.TransactionOutput[i].Mode)
		} else if expectedOutput[i].Receiving != de.TransactionOutput[i].Receiving {
			t.Errorf("Output Receiving (%v) not equal to (%v)", expectedOutput[i].Receiving, de.TransactionOutput[i].Receiving)
		} else if expectedOutput[i].Recipient != de.TransactionOutput[i].Recipient {
			t.Errorf("Output Recipient (%v) not equal to (%v)", expectedOutput[i].Recipient, de.TransactionOutput[i].Recipient)
		} else if expectedOutput[i].Value != de.TransactionOutput[i].Value {
			t.Errorf("Output Value (%v) not equal to (%v)", expectedOutput[i].Value, de.TransactionOutput[i].Value)
		}
	}

	fmt.Printf("%+v\n", de)
}
package models

import (
	"time"
)

type DayExtract struct {
	Date              time.Time     `json:"date"`
	InputValue        float64       `json:"inputValue"`
	TransactionInput  []Transaction `json:"[]transactionInput"`
	OutputValue       float64       `json:"outputValue"`
	TransactionOutput []Transaction `json:"[]transactionOutput"`
}

func (de *DayExtract) GetDayExtract(date time.Time) string {
	// Setting initial date for the day extract
	de.Date = date
	// Fetching the data by a given date
	transactions := FetchAndTransformDailyTransactions()

	// Separate all the transactions into input and output transaction
	for _, transaction := range transactions {
		if transaction.Receiving {
			de.TransactionInput = append(de.TransactionInput, transaction)
		} else {
			de.TransactionOutput = append(de.TransactionOutput, transaction)
		}
	}

	// calculate both the input and output differente within transactions
	de.InputValue = de.calculateInputValue()
	de.OutputValue = de.calculateOutputValue()

	return "status(200)"
}

func (de *DayExtract) calculateInputValue() (outValue float64) {
	for _, transaction := range de.TransactionInput {
		outValue += transaction.Value
	}
	return outValue
}

func (de *DayExtract) calculateOutputValue() (inValue float64) {
	for _, transaction := range de.TransactionOutput {
		inValue += transaction.Value
	}
	return inValue
}

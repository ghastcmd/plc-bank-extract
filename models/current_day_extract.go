package models

import "time"

type DayExtract struct {
	Date              time.Time     `json:"date"`
	InputValue        float64       `json:"inputValue"`
	TransactionInput  []Transaction `json:"[]transactionInput"`
	OutputValue       float64       `json:"outputValue"`
	TransactionOutput []Transaction `json:"[]transactionOutput"`
}

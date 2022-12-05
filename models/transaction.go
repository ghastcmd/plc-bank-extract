package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ghastcmd/plc-bank-extract/database"
)

const (
	Invalid = iota
	CreditCard
	Pix
	Transference
)

type Transaction struct {
	Mode      int     `json:"mode"`
	Receiving bool    `json:"receiving"`
	Recipient string  `json:"recipient"`
	Value     float64 `json:"value"`
}

func (t *Transaction) SetTransaction(
	mode int,
	is_receiving bool,
	recipient string,
	value float64,
) {
	t.Mode = mode
	t.Receiving = is_receiving
	t.Recipient = recipient
	t.Value = value
}

func getTransactionMode(mode string) int {
	switch strings.Title(strings.ToLower(mode)) {
	case "Pix":
		return Pix
	case "Credit Card":
		return CreditCard
	case "Transference":
		return Transference
	default:
		return Invalid
	}
}

func getTransactionReceiving(receiving string) (bool, error) {
	switch strings.ToLower(receiving) {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, errors.New("Invalid token for Receiving")
	}
}

func FetchAndTransformDailyTransactions(date time.Time) (transactions []Transaction) {
	rawTransactions := database.FetchDailyTransactions(date)

	for _, rawTransaction := range rawTransactions {
		mode := getTransactionMode(rawTransaction.Mode)
		receiving, err := getTransactionReceiving(rawTransaction.Receiving)
		if err != nil {
			fmt.Println(err)
		}

		newTransaction := Transaction{
			Mode:      mode,
			Receiving: receiving,
			Recipient: rawTransaction.Recipient,
			Value:     rawTransaction.Value,
		}

		transactions = append(transactions, newTransaction)
	}

	return transactions
}

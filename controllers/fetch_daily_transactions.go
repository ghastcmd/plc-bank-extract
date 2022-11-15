package controllers

import (
	"encoding/json"
	"fmt"
	"time"
)

type JsonRespRaw struct {
	Date      string  `json:"date"`
	Mode      string  `json:"mode"`
	Receiving string  `json:"receiving"`
	Recipient string  `json:"recipient"`
	Value     float64 `json:"value"`
}

var testTransactions = `[
	{
		"date": "2022-09-21",
		"mode": "PIX",
		"receiving": "true",
		"recipient": "378.432.324-89",
		"value": 3232.32
	},
	{
		"date": "2022-09-22",
		"mode": "PIX",
		"receiving": "false",
		"recipient": "378.432.324-89",
		"value": 23.32
	},
	{
		"date": "2022-09-21",
		"mode": "PIX",
		"receiving": "true",
		"recipient": "323.432.324-89",
		"value": 123.32
	},
	{
		"date": "2022-09-23",
		"mode": "PIX",
		"receiving": "true",
		"recipient": "378.431.874-89",
		"value": 1.32
	}
]`

func FetchDailyTransactions(date time.Time) (retTransactions []JsonRespRaw) {
	currentDate := fmt.Sprintf("%v-%02v-%02v", date.Year(), int(date.Month()), date.Day())

	var iterTransactions []JsonRespRaw
	if err := json.Unmarshal([]byte(testTransactions), &iterTransactions); err != nil {
		fmt.Println(err)
	}

	for _, transaction := range iterTransactions {
		if transaction.Date == currentDate {
			retTransactions = append(retTransactions, transaction)
		}
	}

	return retTransactions
}

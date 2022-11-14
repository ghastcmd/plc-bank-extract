package controllers

import (
	"encoding/json"
	"fmt"
)

type JsonRespRaw struct {
	Date      string  `json:"date"`
	Mode      string  `json:"mode"`
	Receiving string  `json:"receiving"`
	Recipient string  `json:"recipient"`
	Value     float64 `json:"value"`
}

func FetchDailyTransactions() []JsonRespRaw {
	test_transactions := `[
	{
		"date": "2022-09-21",
		"mode": "PIX",
		"receiving": "true",
		"recipient": "378.432.324-89",
		"value": 3232.32,
	},
	{
		"date": "2022-09-22",
		"mode": "PIX",
		"receiving": "false",
		"recipient": "378.432.324-89",
		"value": 23.32,
	},
	{
		"date": "2022-09-21",
		"mode": "PIX",
		"receiving": "true",
		"recipient": "323.432.324-89",
		"value": 123.32,
	},
	{
		"date": "2022-09-23",
		"mode": "PIX",
		"receiving": "true",
		"recipient": "378.431.874-89",
		"value": 1.32,
	},
]`

	str, err := json.Marshal(test_transactions)
	if err != nil {
		fmt.Println(err)
	}

	var ret_transactions []JsonRespRaw
	err = json.Unmarshal([]byte(str), &ret_transactions)
	if err != nil {
		fmt.Println(err)
	}

	return ret_transactions
}

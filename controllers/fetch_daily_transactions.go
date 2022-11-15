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

func FetchDailyTransactions() (retTransactions []JsonRespRaw) {
	testTransactions := `[
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

	if err := json.Unmarshal([]byte(testTransactions), &retTransactions); err != nil {
		fmt.Printf("%v %v\n", 56, err)
	}

	return retTransactions
}

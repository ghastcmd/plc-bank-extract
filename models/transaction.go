package models

import "fmt"

type Transaction struct {
	Mode      int     `json:"mode"`
	Recipient string  `json:"recipient"`
	Value     float64 `json:"value"`
}

func HelloFromModel() {
	fmt.Println("Hello from <model> package.")
}

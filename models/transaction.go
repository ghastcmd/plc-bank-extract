package models

type Transaction struct {
	Mode      int     `json:"mode"`
	Recipient string  `json:"recipient"`
	Value     float64 `json:"value"`
}

func (t *Transaction) SetTransaction(mode int, recipient string, value float64) {
	t.Mode = mode
	t.Recipient = recipient
	t.Value = value
}

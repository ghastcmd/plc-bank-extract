package models

const (
	CreditCard = iota
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

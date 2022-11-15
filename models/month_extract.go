package models

import "time"

type MonthExtract struct {
	Month            time.Month   `json:"month"`
	OpeningBalance   float64      `json:"openingBalance"`
	TotalInputs      float64      `json:"totalInputs"`
	TotalOutputs     float64      `json:"totalOutputs"`
	FinalBalance     float64      `json:"totalBalance"`
	ExtractsPerMonth []DayExtract `json:"extractsPerMonth"`
}

func (me *MonthExtract) GetMonthExtract() (s string) {

	return s
}

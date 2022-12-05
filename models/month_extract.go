package models

import (
	"time"
)

type MonthExtract struct {
	Month            time.Month   `json:"month"`
	OpeningBalance   float64      `json:"openingBalance"`
	TotalInputs      float64      `json:"totalInputs"`
	TotalOutputs     float64      `json:"totalOutputs"`
	FinalBalance     float64      `json:"totalBalance"`
	ExtractsPerMonth []DayExtract `json:"extractsPerMonth"`
}

func (me *MonthExtract) GetMonthExtract(date time.Time) (s string) {
	year := date.Year()
	month := date.Month()

	// Assigning month
	me.Month = month

	// Calulating the last day of current month
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, date.Location())
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	// Fetching all transaction data for the entire month
	last := lastOfMonth.Day()
	for i := firstOfMonth.Day(); i <= last; i++ {
		currentDate := time.Date(year, month, i, 0, 0, 0, 0, date.Location())
		var de DayExtract
		de.GetDayExtract(currentDate)
		if de.InputValue != 0 || de.OutputValue != 0 {
			me.ExtractsPerMonth = append(me.ExtractsPerMonth, de)
		}

		// Calculate total inputs
		me.TotalInputs += de.InputValue
		// Calculate total outputs
		me.TotalOutputs += de.OutputValue
	}

	return s
}

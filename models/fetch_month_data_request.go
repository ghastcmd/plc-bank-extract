package models

import "time"

type MonthDataRequest struct {
	Year  int        `json:"year"`
	Month time.Month `json:"month"`
}

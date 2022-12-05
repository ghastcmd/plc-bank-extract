package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ghastcmd/plc-bank-extract/models"
)

func FetchMonthData(w http.ResponseWriter, r *http.Request) {
	var req models.MonthDataRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		badRequest(w, r)
		return
	}

	fmt.Printf("month: %d, year: %d\n", req.Month, req.Year)

	var me models.MonthExtract
	me.GetMonthExtract(time.Date(req.Year, req.Month, 1, 0, 0, 0, 0, time.UTC))

	jsonBytes, err := json.Marshal(me)
	if err != nil {
		internalServerError(w, r)
		return
	}

	w.Write(jsonBytes)
}

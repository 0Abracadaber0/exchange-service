package models

import (
	"time"
)

type ExchangeRate struct {
	CurID           int       `json:"Cur_ID"`
	Date            time.Time `json:"Date"`
	CurAbbreviation string    `json:"Cur_Abbreviation"`
	CurName         string    `json:"Cur_Name"`
	CurScale        int       `json:"Cur_Scale"`
	CurOfficialRate int       `json:"Cur_OfficialRate"`
}

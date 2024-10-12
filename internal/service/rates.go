package service

import (
	"exchange/internal/database"
	model "exchange/internal/models"
	"fmt"
)

func GetAllRates() ([]model.ExchangeRate, error) {
	rows, err := database.DB.Query(`
		SELECT cur_id, date, cur_abbreviation, cur_scale, cur_name, cur_official_rate 
		FROM exchange_rates`)
	if err != nil {
		return []model.ExchangeRate{}, fmt.Errorf("Error fetching data: %v", err)
	}
	defer rows.Close()

	var rates []model.ExchangeRate

	for rows.Next() {
		var rate model.ExchangeRate
		if err := rows.Scan(&rate.CurID, &rate.Date, &rate.CurAbbreviation, &rate.CurScale, &rate.CurName, &rate.CurOfficialRate); err != nil {
			return []model.ExchangeRate{}, fmt.Errorf("Error scanning data: %v", err)
		}
		rates = append(rates, rate)
	}

	if err := rows.Err(); err != nil {
		return []model.ExchangeRate{}, fmt.Errorf("Error iterating data: %v", err)
	}
	return rates, nil
}

package service

import (
	"exchange/internal/database"
	model "exchange/internal/models"
	"fmt"
	"time"
)

func GetAllRates() ([]model.ExchangeRate, error) {
	rows, err := database.DB.Query(`
		SELECT cur_id, date, cur_abbreviation, cur_scale, cur_name, cur_official_rate 
		FROM exchange_rates`)
	if err != nil {
		return []model.ExchangeRate{}, fmt.Errorf("error fetching data: %v", err)
	}
	defer rows.Close()

	var rates []model.ExchangeRate

	for rows.Next() {
		var rate model.ExchangeRate
		if err := rows.Scan(&rate.CurID, &rate.Date, &rate.CurAbbreviation, &rate.CurScale, &rate.CurName, &rate.CurOfficialRate); err != nil {
			return []model.ExchangeRate{}, fmt.Errorf("error scanning data: %v", err)
		}
		rates = append(rates, rate)
	}

	if err := rows.Err(); err != nil {
		return []model.ExchangeRate{}, fmt.Errorf("error iterating data: %v", err)
	}
	return rates, nil
}

func GetRatesByDate(date string) ([]model.ExchangeRate, error) {
	formattedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %w", err)
	}

	rows, err := database.DB.Query(`
		SELECT cur_id, date, cur_abbreviation, cur_scale, cur_name, cur_official_rate 
		FROM exchange_rates
		WHERE date = ?`, formattedDate)
	if err != nil {
		return []model.ExchangeRate{}, fmt.Errorf("error fetching data: %v", err)
	}
	defer rows.Close()

	var rates []model.ExchangeRate

	for rows.Next() {
		var rate model.ExchangeRate
		if err := rows.Scan(&rate.CurID, &rate.Date, &rate.CurAbbreviation, &rate.CurScale, &rate.CurName, &rate.CurOfficialRate); err != nil {
			return []model.ExchangeRate{}, fmt.Errorf("error scanning data: %v", err)
		}
		rates = append(rates, rate)
	}

	if err := rows.Err(); err != nil {
		return []model.ExchangeRate{}, fmt.Errorf("error iterating data: %v", err)
	}
	return rates, nil
}

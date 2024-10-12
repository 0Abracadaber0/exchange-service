package service

import (
	"encoding/json"
	"exchange/internal/config"
	"exchange/internal/database"
	model "exchange/internal/models"
	"fmt"
	"io"
	"net/http"
	"time"

	"log/slog"
)

func FetchAndStoreData(log *slog.Logger, cfg *config.Config) error {
	log.Debug("FetchAndStoreData called")
	uri := "https://api.nbrb.by/exrates/rates?periodicity=0"

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(uri)
	if err != nil {
		return fmt.Errorf("error sending the request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed to read response body: ", "err", err)
		return fmt.Errorf("error reading response body: %w", err)
	}

	var rates []model.ExchangeRate

	err = json.Unmarshal(body, &rates)
	if err != nil {
		log.Error("Failed to decode JSON", "error", err)
		return fmt.Errorf("error decoding JSON: %w", err)
	}

	log.Debug("rates: ", "rates", rates)
	for _, rate := range rates {
		parsedDate, err := time.Parse("2006-01-02T15:04:05", rate.Date)
		if err != nil {
			log.Error("Failed to parse date", "date", rate.Date, "error", err)
			continue
		}

		_, err = database.DB.Exec(`
			INSERT INTO exchange_rates (cur_id, date, cur_abbreviation, cur_scale, cur_name, cur_official_rate)
			VALUES (?, ?, ?, ?, ?, ?)
			ON DUPLICATE KEY UPDATE
			cur_official_rate = VALUES(cur_official_rate)`,
			rate.CurID, parsedDate, rate.CurAbbreviation, rate.CurScale, rate.CurName, rate.CurOfficialRate)

		if err != nil {
			log.Error("Failed to insert/update data in DB", "currency", rate.CurAbbreviation, "error", err)
		}
	}

	log.Info("successfully fetched and stored currency rates.")

	return nil
}

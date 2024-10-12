package service

import (
	"exchange/internal/database"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllRates(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	database.DB = db

	mock.ExpectQuery(`SELECT cur_id, date, cur_abbreviation, cur_scale, cur_name, cur_official_rate FROM exchange_rates`).
		WillReturnRows(sqlmock.NewRows([]string{"cur_id", "date", "cur_abbreviation", "cur_scale", "cur_name", "cur_official_rate"}).
			AddRow(1, "2024-10-12", "USD", 1, "Доллар США", 3.3).
			AddRow(2, "2024-10-12", "EUR", 1, "Евро", 3.5))

	rates, err := GetAllRates()

	assert.NoError(t, err)
	assert.Len(t, rates, 2)
	assert.Equal(t, rates[0].CurID, 1)
	assert.Equal(t, rates[1].CurAbbreviation, "EUR")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetRatesByDate(t *testing.T) {
	date := "2024-10-12"
	formattedDate, _ := time.Parse("2006-01-02", date)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	database.DB = db

	mock.ExpectQuery(`SELECT cur_id, date, cur_abbreviation, cur_scale, cur_name, cur_official_rate FROM exchange_rates WHERE date = ?`).
		WithArgs(formattedDate).
		WillReturnRows(sqlmock.NewRows([]string{"cur_id", "date", "cur_abbreviation", "cur_scale", "cur_name", "cur_official_rate"}).
			AddRow(1, formattedDate, "USD", 1, "Доллар США", 3.3))

	rates, err := GetRatesByDate(date)

	assert.NoError(t, err)
	assert.Len(t, rates, 1)
	assert.Equal(t, rates[0].CurID, 1)
	assert.Equal(t, rates[0].CurAbbreviation, "USD")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetRatesByDate_InvalidDateFormat(t *testing.T) {
	_, err := GetRatesByDate("invalid-date")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid date format")
}

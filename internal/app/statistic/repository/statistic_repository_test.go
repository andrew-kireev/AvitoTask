package repository

import (
	"AvitoTask/internal/app/statistic/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetStatistic(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	statRep := NewStatisticRepository(db)

	defer db.Close()

	expectedStat := models.Statistic{
		Date:   "2005-04-13",
		Views:  1000,
		Clicks: 200,
		Cost:   decimal.NewFromFloat(0),
		CPC:    decimal.NewFromFloat(0),
		CPM:    decimal.NewFromFloat(0),
	}

	rows := sqlmock.NewRows([]string{
		"stat_date", "views", "clicks", "cost", "cpc", "cpm",
	}).AddRow(expectedStat.Date, expectedStat.Views, expectedStat.Clicks, expectedStat.Cost, 0, 0)
	query := "SELECT"

	mock.ExpectQuery(query).WithArgs("2003-04-13", "2010-04-13").WillReturnRows(rows)

	stats, err := statRep.GetStatistic("2003-04-13", "2010-04-13", "stat_date")

	assert.NoError(t, err)
	if !reflect.DeepEqual(expectedStat, stats[0]) {
		t.Fatalf("Not equal, expected: %v, got: %v", expectedStat, stats[0])
	}
}

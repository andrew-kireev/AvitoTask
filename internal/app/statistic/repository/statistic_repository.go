package repository

import (
	"AvitoTask/internal/app/statistic"
	"AvitoTask/internal/app/statistic/models"
	"database/sql"
	"github.com/shopspring/decimal"
	"strings"
)

type StatisticRepository struct {
	con *sql.DB
}

func NewStatisticRepository(con *sql.DB) statistic.Repository {
	return &StatisticRepository{
		con: con,
	}
}

func (statRep *StatisticRepository) SaveStatistic(stat models.Statistic) error {
	queryGetStat := `SELECT stat_date, views, clicks, cost FROM statistic WHERE stat_date = $1`

	statistic := models.Statistic{}
	err := statRep.con.QueryRow(
		queryGetStat, stat.Date).Scan(&statistic.Date, &statistic.Views, &statistic.Clicks,
		&statistic.Cost)
	if err == nil {
		stat.Views += statistic.Views
		stat.Cost = stat.Cost.Add(statistic.Cost)
		stat.Clicks += statistic.Clicks

		stat.CPC = stat.Cost.Div(decimal.NewFromInt(int64(stat.Clicks)))
		stat.CPM = stat.Cost.Div(decimal.NewFromInt(int64(stat.Views))).Mul(decimal.NewFromInt(1000))

		queryUpdate := `UPDATE statistic SET views = $1, clicks = $2,
                     cost = $3, cpc = $4, cpm = $5 WHERE stat_date = $6`

		_, err = statRep.con.Exec(
			queryUpdate, stat.Views, stat.Clicks, stat.Cost, stat.CPC, stat.CPM, stat.Date)

		if err != nil {
			return err
		}
		return nil
	}
	query := `INSERT INTO statistic (stat_date, views, clicks, cost, cpc, cpm) VALUES
        ($1, $2, $3, $4, $5, $6)`

	stat.CPC = stat.Cost.Div(decimal.NewFromInt(int64(stat.Clicks)))
	stat.CPM = stat.Cost.Div(decimal.NewFromInt(int64(stat.Views))).Mul(decimal.NewFromInt(1000))

	err = statRep.con.QueryRow(
		query, stat.Date, stat.Views, stat.Clicks, stat.Cost, stat.CPC, stat.CPM).Scan()

	if err != nil {
		return err
	}
	return nil
}

func (statRep *StatisticRepository) GetStatistic(from, to, sortingParam string) ([]models.Statistic, error) {
	query := `SELECT stat_date, views, clicks, cost, cpc, cpm FROM statistic
			where stat_date >= $1 and stat_date <= $2
			order by ` + sortingParam

	rows, err := statRep.con.Query(
		query, from, to)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	statistics := make([]models.Statistic, 0)

	for rows.Next() {
		stat := models.Statistic{}
		_ = rows.Scan(&stat.Date, &stat.Views, &stat.Clicks,
			&stat.Cost, &stat.CPC, &stat.CPM)
		stat.Date = strings.Split(stat.Date, "T")[0]
		statistics = append(statistics, stat)
	}

	return statistics, err
}

func (statRep *StatisticRepository) DeleteAll() error {
	queryDelete := `DELETE FROM statistic`

	_, err := statRep.con.Exec(
		queryDelete)
	if err != nil {
		return err
	}
	return nil
}

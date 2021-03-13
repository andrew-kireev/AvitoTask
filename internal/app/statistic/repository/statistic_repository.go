package repository

import (
	"AvitoTask/internal/app/statistic"
	"AvitoTask/internal/app/statistic/models"
	"database/sql"
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
	query := `SELECT stat_date, views, clicks, cost FROM statistic WHERE stat_date = $1`

	rows, err := statRep.con.Query(
		query, stat.Date)
}

package statistic

import "AvitoTask/internal/app/statistic/models"

type Repository interface {
	SaveStatistic(*models.Statistic) error
	DeleteAll() error
	GetStatistic(string, string, string) ([]models.Statistic, error)
}

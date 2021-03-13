package usecase

import (
	"AvitoTask/internal/app/statistic"
	"AvitoTask/internal/app/statistic/models"
)

type StatisticUsecase struct {
	statisticRep statistic.Repository
}

func NewStatisticUsecase(statisticRep statistic.Repository) *StatisticUsecase {
	return &StatisticUsecase{
		statisticRep: statisticRep,
	}
}

func (statUsecase *StatisticUsecase) SaveStatistic(stat *models.Statistic) error {
	err := statUsecase.statisticRep.SaveStatistic(stat)
	if err != nil {
		return err
	}
	return nil
}

func (statUsecase *StatisticUsecase) DeleteAll() error {
	err := statUsecase.statisticRep.DeleteAll()
	if err != nil {
		return err
	}
	return nil
}

func (statUsecase *StatisticUsecase) GetStatistic(from, to, sortingParam string) ([]models.Statistic, error) {
	if sortingParam == "" {
		sortingParam = "stat_date"
	}

	stats, err := statUsecase.statisticRep.GetStatistic(from, to, sortingParam)
	if err != nil {
		return nil, err
	}
	return stats, nil
}

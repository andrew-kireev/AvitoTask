package usecase

import "AvitoTask/internal/app/statistic"

type StatisticUsecase struct {
	statisticRep statistic.Repository
}

func NewStatisticUsecase(statisticRep statistic.Repository) *StatisticUsecase {
	return &StatisticUsecase{
		statisticRep: statisticRep,
	}
}



package usecase

import (
	mock_statistic "AvitoTask/internal/app/statistic/mocks"
	"AvitoTask/internal/app/statistic/models"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveStatistic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_statistic.NewMockRepository(ctrl)
	mockUsecase := NewStatisticUsecase(mockRepo)

	expectedStatistic := &models.Statistic{
		Date:   "2010-07-13",
		Views:  50,
		Clicks: 200,
		Cost:   decimal.NewFromFloat(1000.30),
	}

	mockRepo.
		EXPECT().SaveStatistic(gomock.Eq(expectedStatistic)).
		Return(nil)

	err := mockUsecase.SaveStatistic(expectedStatistic)
	assert.Equal(t, err, nil)
}

func TestGetStatistic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_statistic.NewMockRepository(ctrl)
	mockUsecase := NewStatisticUsecase(mockRepo)

	expectedStatistic := []models.Statistic{
		{
			Date:   "2010-07-13",
			Views:  50,
			Clicks: 200,
			Cost:   decimal.NewFromFloat(1000.30),
		},
	}

	mockRepo.
		EXPECT().GetStatistic("2006-07-13", "2009-07-13", "cpm").
		Return(expectedStatistic, nil)

	some, err := mockUsecase.GetStatistic("2006-07-13", "2009-07-13", "cpm")
	fmt.Println(some)
	assert.Equal(t, err, nil)
	assert.Equal(t, some, expectedStatistic)
}

func TestDeleteAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_statistic.NewMockRepository(ctrl)
	mockUsecase := NewStatisticUsecase(mockRepo)

	mockRepo.
		EXPECT().DeleteAll().
		Return(nil)

	err := mockUsecase.DeleteAll()
	assert.Equal(t, err, nil)
}

package http

import (
	"AvitoTask/configs"
	mock_statistic "AvitoTask/internal/app/statistic/mocks"
	"AvitoTask/internal/app/statistic/models"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var (
	statistics = []models.Statistic{
		{
			Date:   "2010-07-13",
			Views:  50,
			Clicks: 200,
			Cost:   decimal.NewFromFloat(1000.30),
		},
		{
			Date:   "2005-07-13",
			Views:  100,
			Clicks: 300,
			Cost:   decimal.NewFromFloat(1000.30),
		},
	}
)

func TestDeleteStatisticHandlerSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStatisticUsecase := mock_statistic.NewMockUsecase(ctrl)

	mockStatisticUsecase.EXPECT().DeleteAll().Times(1).Return(nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/api/vi/statistic/", nil)

	handler := NewStatisticHandler(mux.NewRouter(), configs.NewConfig(), mockStatisticUsecase)

	handler.DeleteStatisticHandler(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	if !reflect.DeepEqual("", w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", "", w.Body.String())
	}
}

func TestDeleteStatisticHandlerFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStatisticUsecase := mock_statistic.NewMockUsecase(ctrl)

	mockStatisticUsecase.EXPECT().DeleteAll().Times(1).Return(errors.New("some error"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/api/vi/statistic/", nil)

	handler := NewStatisticHandler(mux.NewRouter(), configs.NewConfig(), mockStatisticUsecase)

	handler.DeleteStatisticHandler(w, r)

	expected := http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestGetStatisticHandlerSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStatisticUsecase := mock_statistic.NewMockUsecase(ctrl)

	mockStatisticUsecase.EXPECT().GetStatistic("2006-07-13",
		"2009-07-13", "stat_date").Times(1).Return(statistics, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET",
		"/api/vi/statistic?from=2006-07-13&&to=2009-07-13&&sort=stat_date", nil)

	handler := NewStatisticHandler(mux.NewRouter(), configs.NewConfig(), mockStatisticUsecase)

	handler.GetStatisticHandler(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(statistics)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}
}

func TestGetStatisticHandlerFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStatisticUsecase := mock_statistic.NewMockUsecase(ctrl)

	mockStatisticUsecase.EXPECT().GetStatistic("2006-07-13",
		"2009-07-13", "stat_date").Times(1).Return(nil, errors.New("some error"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET",
		"/api/vi/statistic?from=2006-07-13&&to=2009-07-13&&sort=stat_date", nil)

	handler := NewStatisticHandler(mux.NewRouter(), configs.NewConfig(), mockStatisticUsecase)

	handler.GetStatisticHandler(w, r)

	expected := http.StatusNoContent
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	if !reflect.DeepEqual("", w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", "", w.Body.String())
	}
}

func TestSaveStatisticHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStatisticUsecase := mock_statistic.NewMockUsecase(ctrl)

	mockStatisticUsecase.EXPECT().SaveStatistic(&statistics[0]).Times(1).Return(nil)

	w := httptest.NewRecorder()

	stat := models.Request{
		Date:   "2010-07-13",
		Views:  "50",
		Clicks: "200",
		Cost:   "1000.3",
	}

	jsonBody, _ := json.Marshal(&stat)
	body := bytes.NewReader(jsonBody)

	r := httptest.NewRequest("POST",
		"/api/vi/statistic", body)

	handler := NewStatisticHandler(mux.NewRouter(), configs.NewConfig(), mockStatisticUsecase)

	handler.SaveStatisticHandler(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestSaveStatisticHandlerFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStatisticUsecase := mock_statistic.NewMockUsecase(ctrl)
	mockStatisticUsecase.EXPECT().SaveStatistic(&statistics[0]).Times(1).Return(errors.New("some error"))

	w := httptest.NewRecorder()

	stat := models.Request{
		Date:   "2010-07-13",
		Views:  "50",
		Clicks: "200",
		Cost:   "1000.3",
	}

	jsonBody, _ := json.Marshal(&stat)
	body := bytes.NewReader(jsonBody)

	r := httptest.NewRequest("POST",
		"/api/vi/statistic", body)

	handler := NewStatisticHandler(mux.NewRouter(), configs.NewConfig(), mockStatisticUsecase)

	handler.SaveStatisticHandler(w, r)

	expected := http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

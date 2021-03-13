package http

import (
	"AvitoTask/configs"
	"AvitoTask/internal/app/middlewares"
	"AvitoTask/internal/app/statistic"
	"AvitoTask/internal/app/statistic/models"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
)

type StatisticHandler struct {
	router      *mux.Router
	logger      *logrus.Logger
	statUsecase statistic.Usecase
	validator   *validator.Validate
}

func NewStatisticHandler(r *mux.Router, config *configs.Config, usecase statistic.Usecase) *StatisticHandler {
	handler := &StatisticHandler{
		router:      r,
		logger:      logrus.New(),
		statUsecase: usecase,
	}

	err := ConfigLogger(handler, config)
	if err != nil {
		logrus.Error(err)
	}

	handler.validator = validator.New()

	handler.router.HandleFunc("/api/v1/statistic",
		handler.SaveStatisticHandler).Methods(http.MethodPost)
	handler.router.HandleFunc("/api/v1/statistic",
		handler.DeleteStatisticHandler).Methods(http.MethodDelete)
	handler.router.HandleFunc("/api/v1/statistic",
		handler.GetStatisticHandler).Methods(http.MethodGet)

	handler.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("main of albums"))
	})

	handler.router.Use(middlewares.LoggingMiddleware)
	handler.router.Use(middlewares.PanicMiddleware)
	handler.router.Use(middlewares.ContentTypeJson)

	return handler
}

func (handler *StatisticHandler) SaveStatisticHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.Request{}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handler.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := req.UnmarshalJSON(data); err != nil {
		handler.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = handler.validator.Struct(req)
	if err != nil {
		handler.logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	stat := &models.Statistic{}
	stat.Date = req.Date
	stat.Views, _ = strconv.Atoi(req.Views)
	stat.Cost, _ = decimal.NewFromString(req.Cost)
	stat.Clicks, _ = strconv.Atoi(req.Clicks)

	err = handler.statUsecase.SaveStatistic(stat)
	if err != nil {
		handler.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *StatisticHandler) DeleteStatisticHandler(w http.ResponseWriter, r *http.Request) {
	err := handler.statUsecase.DeleteAll()
	if err != nil {
		handler.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *StatisticHandler) GetStatisticHandler(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	sortingParam := r.URL.Query().Get("sort")
	stats, err := handler.statUsecase.GetStatistic(from, to, sortingParam)
	if err != nil {
		handler.logger.Error(err)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	body, err := json.Marshal(stats)
	if err != nil {
		handler.logger.Error(err)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (handler *StatisticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}

func ConfigLogger(handler *StatisticHandler, config *configs.Config) error {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}

	handler.logger.SetLevel(level)
	return nil
}

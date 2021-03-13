package http

import (
	"AvitoTask/configs"
	"AvitoTask/internal/app/statistic"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type StatisticHandler struct {
	router           *mux.Router
	logger           *logrus.Logger
	statisticUsecase statistic.Usecase
}

func NewStatisticHandler(r *mux.Router, config *configs.Config, usecase statistic.Usecase) *StatisticHandler {
	handler := &StatisticHandler{
		router:           r,
		logger:           logrus.New(),
		statisticUsecase: usecase,
	}

	err := ConfigLogger(handler, config)
	if err != nil {
		logrus.Error(err)
	}

	//handler.router.HandleFunc("/{album_id:[0-9]+}", handler).Methods("GET")

	handler.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("main of albums"))
	})

	return handler
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

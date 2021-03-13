package main

import (
	"AvitoTask/configs"
	"AvitoTask/internal/app/statistic/delivery/http"
	"AvitoTask/internal/app/statistic/repository"
	"AvitoTask/internal/app/statistic/usecase"
	"AvitoTask/internal/pkg/server"
	"AvitoTask/internal/pkg/utility"
	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	configsPath = "configs/config.toml"
)

func main() {
	//time.Sleep(6 * time.Second)
	config := configs.NewConfig()
	_, err := toml.DecodeFile(configsPath, config)
	if err != nil {
		logrus.Error(err)
	}

	postgresCon, err := utility.CreatePostgresConnection(config.StatisticPostgresBD)
	if err != nil {
		logrus.Error(err)
		return
	}
	router := mux.NewRouter()

	statRep := repository.NewStatisticRepository(postgresCon)
	statUsecase := usecase.NewStatisticUsecase(statRep)

	statHandler := http.NewStatisticHandler(router, config, statUsecase)

	if err = server.Start(config, statHandler); err != nil {
		logrus.Error(err)
	}
}

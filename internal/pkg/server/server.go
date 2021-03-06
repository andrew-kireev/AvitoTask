package server

import (
	"AvitoTask/configs"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	handler Handler
	config  *configs.Config
}

func NewServer(config *configs.Config, handler Handler) (*Server, error) {
	serv := &Server{
		config:  config,
		handler: handler,
	}

	return serv, nil
}

func Start(config *configs.Config, handler Handler) error {
	serv, err := NewServer(config, handler)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return http.ListenAndServe(serv.config.StatisticServerAddr, serv.handler)
}

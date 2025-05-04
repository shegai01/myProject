package api

import (
	"go.uber.org/zap"
)

type API struct {
	config   *Config
	logLevel *zap.SugaredLogger
}

func NewAPI(config *Config) *API {
	return &API{
		config:   config,
		logLevel: &zap.SugaredLogger{},
	}

}

package api

import (
	"net/http"

	"github.com/shegai01/myProject/internal/router"
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
func (api *API) Start(log zap.SugaredLogger) error {
	// log.Info("server starting port:8080")
	return (http.ListenAndServe(":"+api.config.BindAddr, router.CustomRouter()))
	// return nil
}

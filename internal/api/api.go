package api

import (
	"net/http"

	"github.com/gorilla/mux"
	handler "github.com/shegai01/myProject/internal/handlers"
	zaplog "github.com/shegai01/myProject/internal/logger"
	"github.com/shegai01/myProject/internal/router"

	"go.uber.org/zap"
)

type API struct {
	config *Config
	logger *zap.SugaredLogger
	router *mux.Router
}

func NewAPI(config *Config) *API {
	return &API{
		config: config,
		logger: zaplog.ZapLog(),
		router: router.CustomRouter(),
	}

}
func (a *API) Start() error {
	// a.configLogger()
	// a.configRouter()
	a.logger.Info("server starting port:8081")
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("hello router is works\n"))
	})
	group := a.router.PathPrefix("/user").Subrouter()
	group.HandleFunc("/create", handler.CreateUser)
	a.logger.Debug("listen")
	if err := http.ListenAndServe(":"+a.config.BindAddr, a.router); err != nil {
		a.logger.Infoln("connection failed", err)
		return err
	}
	return nil
}

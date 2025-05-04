package main

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
	"github.com/shegai01/myProject/internal/api"
	"github.com/shegai01/myProject/internal/logger"
	"go.uber.org/zap"
)

var (
	configToml string
	configENV  string
)

func init() {
	flag.StringVar(&configToml, "toml", "config/api.toml", "not finded file")
	flag.StringVar(&configENV, "env", "config/api.env", "not finded file")
}
func main() {
	flag.Parse()

	log, err := logger.ZapLog()
	if err != nil {
		return
	}
	err = godotenv.Load(configENV)
	if err != nil {
		return
	}
	portENV := os.Getenv("bind_addr")
	log.Info("server starting at port: %s", portENV)
	config := api.Config{BindAddr: portENV}
	server := api.NewAPI(&config)
	if err := server.Start(*log); err != nil {
		log.Error("connection failed", zap.Error(err))
		return
	}
}

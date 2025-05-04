package main

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
	"github.com/shegai01/myProject/internal/api"
)

var (
	configToml string
	configENV  string
)

func init() {
	// flag.StringVar(&configToml, "toml", "config/api.toml", "not finded file")
	flag.StringVar(&configENV, "env", "config/api.env", "not finded file")
}
func main() {
	flag.Parse()

	err := godotenv.Load(configENV)
	if err != nil {
		return
	}

	portENV := os.Getenv("bind_addr")
	// log.Info("server starting at port: %s", portENV)
	config := api.Config{BindAddr: portENV}
	server := api.NewAPI(&config)
	if err := server.Start(); err != nil {
		return
	}
}

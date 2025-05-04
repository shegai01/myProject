package main

import (
	"flag"

	"github.com/shegai01/myProject/internal/logger"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "config/api.toml", "not finded file")
}
func main() {
	flag.Parse()
	log, err := logger.ZapLog()
	if err != nil {
		return
	}
	log.Info("server starting")
}

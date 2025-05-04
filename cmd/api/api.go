package main

import (
	"flag"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "config/api.toml", "not finded file")
}
func main() {
	flag.Parse()

	log.Println("server starting")
}

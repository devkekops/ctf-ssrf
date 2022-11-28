package main

import (
	"flag"

	"github.com/caarlos0/env"
	"github.com/devkekops/ctf-ssrf/internal/app/config"
	"github.com/devkekops/ctf-ssrf/internal/app/logger"
	"github.com/devkekops/ctf-ssrf/internal/app/server"
)

func main() {
	logger.InitLog()

	var cfg config.Config
	err := env.Parse(&cfg)
	if err != nil {
		logger.Logger.Fatal().Err(err).Msg("")
		return
	}

	flag.StringVar(&cfg.ServerAddress, "a", cfg.ServerAddress, "server address")
	flag.StringVar(&cfg.Flag, "f", cfg.Flag, "flag")
	flag.Parse()

	logger.Logger.Fatal().Err(server.Serve(&cfg)).Msg("")
}

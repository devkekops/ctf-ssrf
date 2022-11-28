package server

import (
	"net/http"

	"github.com/devkekops/ctf-ssrf/internal/app/client"
	"github.com/devkekops/ctf-ssrf/internal/app/config"
	"github.com/devkekops/ctf-ssrf/internal/app/handlers"
)

func Serve(cfg *config.Config) error {
	client := client.NewCli()

	var baseHandler = handlers.NewBaseHandler(client, cfg.Flag)

	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: baseHandler,
	}

	return server.ListenAndServe()
}

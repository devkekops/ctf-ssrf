package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/devkekops/ctf-ssrf/internal/app/client"
	"github.com/devkekops/ctf-ssrf/internal/app/logger"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type BaseHandler struct {
	mux    *chi.Mux
	fs     http.Handler
	client client.Client
}

var (
	opsProcessed = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "ctf_log_file_size",
		Help: "Size of log file. Hello ctf hacker :) you need this file, check contents of /tmp/log.txt",
	})
)

func recordMetrics() {
	go func() {
		for {
			fi, err := os.Stat("/tmp/log.txt")
			if err != nil {
				logger.Logger.Err(err).Msg("")
			}
			opsProcessed.Set(float64(fi.Size()))
			time.Sleep(2 * time.Second)
		}
	}()
}

func NewBaseHandler(client client.Client, flag string) *chi.Mux {
	logger.Logger.Info().Msg("hello again, ctf hacker :) this is your flag: " + flag)

	cwd, _ := os.Getwd()
	root := filepath.Join(cwd, "/static")
	fs := http.FileServer(http.Dir(root))

	bh := &BaseHandler{
		mux:    chi.NewMux(),
		fs:     fs,
		client: client,
	}

	recordMetrics()

	bh.mux.Handle("/*", fs)
	bh.mux.Get("/convert", bh.convert())
	bh.mux.Handle("/metrics", promhttp.Handler())

	return bh.mux
}

func (bh *BaseHandler) convert() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		logger.Logger.Info().Msg(req.Method + " " + req.RequestURI)

		url := req.URL.Query().Get("url")

		out, err := bh.client.GetInfo(url)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Logger.Err(err).Msg("")
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, err = w.Write([]byte(out))
		if err != nil {
			logger.Logger.Err(err).Msg("")
		}
	}
}

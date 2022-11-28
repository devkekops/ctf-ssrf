package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/middleware"
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

func NewBaseHandler(client client.Client, flag string) *chi.Mux {
	logger.Logger.Info().Msg("hello, hacker) this is your flag: " + flag)

	cwd, _ := os.Getwd()
	root := filepath.Join(cwd, "/static")
	fs := http.FileServer(http.Dir(root))

	bh := &BaseHandler{
		mux:    chi.NewMux(),
		fs:     fs,
		client: client,
	}

	recordMetrics()

	bh.mux.Use(middleware.Logger)

	bh.mux.Handle("/*", fs)
	bh.mux.Get("/convert", bh.convert())
	bh.mux.Handle("/metrics", promhttp.Handler())

	return bh.mux
}

func (bh *BaseHandler) convert() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
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

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

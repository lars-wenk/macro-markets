package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"

	"github.com/lars-wenk/macro-markets/internal/config"
	"github.com/lars-wenk/macro-markets/internal/db"
	"github.com/lars-wenk/macro-markets/internal/health"
	httpapi "github.com/lars-wenk/macro-markets/internal/http"

	"github.com/lars-wenk/macro-markets/internal/logger"
	"github.com/lars-wenk/macro-markets/internal/queue"
)

func main() {
	logger.Setup()
	cfg := config.Load()

	ctx := context.Background()
	pool := db.MustConnect(ctx, cfg)
	defer pool.Close()

	nc := queue.MustConnect(cfg)
	defer nc.Drain()

	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Mount("/healthz", health.Router(pool, nc))
	r.Mount("/api/v1", httpapi.Router(pool, nc, cfg))

	srv := &http.Server{Addr: cfg.HTTPAddr, Handler: r, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second}
	log.Info().Msgf("api-gateway listening on %s", cfg.HTTPAddr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("server error")
	}
}

package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lars-wenk/macro-markets/internal/config"
	"github.com/rs/zerolog/log"
)

func MustConnect(ctx context.Context, cfg config.Config) *pgxpool.Pool {
	url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", cfg.PGUser, cfg.PGPass, cfg.PGHost, cfg.PGPort, cfg.PGDB, cfg.PGSSLMode)
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		log.Fatal().Err(err).Msg("pgxpool new")
	}
	if err := pool.Ping(ctx); err != nil {
		log.Fatal().Err(err).Msg("pg ping")
	}
	log.Info().Msg("postgres connected")
	return pool
}

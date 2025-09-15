package queue

import (
	"github.com/lars-wenk/macro-markets/internal/config"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

func MustConnect(cfg config.Config) *nats.Conn {
	nc, err := nats.Connect(cfg.NatsURL)
	if err != nil {
		log.Fatal().Err(err).Msg("nats connect")
	}
	log.Info().Msg("nats connected")
	return nc
}

func MustConnectWithJS(url string) (*nats.Conn, nats.JetStreamContext) {
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal().Err(err).Msg("nats connect")
	}
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal().Err(err).Msg("jetstream")
	}
	return nc, js
}

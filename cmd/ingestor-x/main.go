package main


import (
"context"
"encoding/json"
"os"
"time"


"github.com/nats-io/nats.go"
"github.com/rs/zerolog/log"


"github.com/yourorg/macro-markets/internal/logger"
"github.com/yourorg/macro-markets/internal/queue"
"github.com/yourorg/macro-markets/pkg/models"
)


func main() {
logger.Setup()
cfg := struct{ NatsURL, Subject string }{os.Getenv("NATS_URL"), os.Getenv("NATS_SUBJECT")}
if cfg.NatsURL == "" { cfg.NatsURL = "nats://localhost:4222" }
if cfg.Subject == "" { cfg.Subject = "ingest.items" }


nc, js := queue.MustConnectWithJS(cfg.NatsURL)
defer nc.Drain()


log.Info().Msgf("ingestor-x publishing demo items to %s", cfg.Subject)
for {
item := models.IngestItem{SourceID: "demo", Kind: "tweet", Data: json.RawMessage(`{"text":"Hello, markets!","author":"@demo"}`)}
b, _ := json.Marshal(item)
if _, err := js.Publish(cfg.Subject, b); err != nil {
log.Error().Err(err).Msg("publish failed")
}
time.Sleep(5 * time.Second)
}
}
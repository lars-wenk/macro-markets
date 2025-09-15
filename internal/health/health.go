package health


import (
"net/http"


"github.com/go-chi/chi/v5"
"github.com/jackc/pgx/v5/pgxpool"
"github.com/nats-io/nats.go"
)


func Router(pool *pgxpool.Pool, nc *nats.Conn) chi.Router {
	r := chi.NewRouter() 
	r.Get("/live", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(200) })
	r.Get("/ready", func(w http.ResponseWriter, _ *http.Request) { 
	if pool == nil || nc == nil { http.Error(w, "deps not ready", 503); return }
		w.WriteHeader(200)
	})
	return r
}
package httpapi


import (
"encoding/json"
"net/http"
"time"


"github.com/go-chi/chi/v5"
"github.com/go-chi/chi/v5/middleware"
"github.com/jackc/pgx/v5/pgxpool"
"github.com/nats-io/nats.go"
"github.com/yourorg/macro-markets/internal/auth"
"github.com/yourorg/macro-markets/internal/config"
)


func Router(pool *pgxpool.Pool, nc *nats.Conn, cfg config.Config) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger) 
	r.Route("/", func(r chi.Router) {
	r.Get("/time", func(w http.ResponseWriter, _ *http.Request) { json.NewEncoder(w).Encode(map[string]string{"now": time.Now().UTC().Format(time.RFC3339)}) })
	r.Mount("/briefings", briefingsRoutes())
	r.Mount("/calendar", calendarRoutes())
	})  
	// Example protected group
	r.Group(func(pr chi.Router) { 
	pr.Use(auth.JWTMiddleware(cfg.JWTSecret)) 
	pr.Get("/me", func(w http.ResponseWriter, _ *http.Request) { json.NewEncoder(w).Encode(map[string]string{"role": "pro"}) })
	}) 
	return r 
}


func briefingsRoutes() chi.Router {
	r := chi.NewRouter()   
	r.Get("/latest", func(w http.ResponseWriter, _ *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]any{"type": "am", "items": []string{"Gold stabil", "DXY +0.2%"}})
	})
	return r 
}


func calendarRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]any{"events": []map[string]string{{"name": "US CPI", "when": "2025-10-10T12:30:00Z"}}})
	})
	return r
}
package config


import "os"


type Config struct {
HTTPAddr string
PGHost string
PGPort string
PGUser string
PGPass string
PGDB string
PGSSLMode string
NatsURL string
JWTSecret string
}


func Load() Config {
cfg := Config{
HTTPAddr: getEnv("HTTP_ADDR", ":8080"),
PGHost: getEnv("PG_HOST", "localhost"),
PGPort: getEnv("PG_PORT", "5432"),
PGUser: getEnv("PG_USER", "mm"),
PGPass: getEnv("PG_PASSWORD", "mm_pw"),
PGDB: getEnv("PG_DATABASE", "mmdb"),
PGSSLMode: getEnv("PG_SSLMODE", "disable"),
NatsURL: getEnv("NATS_URL", "nats://localhost:4222"),
JWTSecret: getEnv("JWT_SECRET", "dev_secret_change_me"),
}
return cfg
}


func getEnv(k, def string) string { if v := os.Getenv(k); v != "" { return v }; return def }

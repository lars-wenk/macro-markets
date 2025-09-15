SHELL := /bin/bash


.PHONY: dev run build test lint up down logs

 
up:
docker compose up -d db nats
sleep 2
docker compose ps


DOWN_SERVICES := db nats api ingestor_x


down:
docker compose down


run:
go run ./cmd/api-gateway


build:
go build -o bin/api ./cmd/api-gateway
go build -o bin/ingestor-x ./cmd/ingestor-x


test:
go test ./...


lint:
@echo "(add golangci-lint if desired)"


migrate:
psql "postgresql://$${PG_USER:-mm}:$${PG_PASSWORD:-mm_pw}@$${PG_HOST:-localhost}:$${PG_PORT:-5432}/$${PG_DATABASE:-mmdb}?sslmode=$${PG_SSLMODE:-disable}" -f migrations/001_init.sql
```
```
# For macOS without `psql`, install with: brew install libpq && brew link --force libpq
```
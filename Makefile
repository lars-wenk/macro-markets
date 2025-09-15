SHELL := /bin/bash

.PHONY: up down run build test lint migrate

up:
	 docker compose up -d db nats
	 sleep 2
	 docker compose ps

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
	@if command -v psql >/dev/null 2>&1; then \
	  psql "postgresql://$${PG_USER:-mm}:$${PG_PASSWORD:-mm_pw}@$${PG_HOST:-localhost}:$${PG_PORT:-5432}/$${PG_DATABASE:-mmdb}?sslmode=$${PG_SSLMODE:-disable}" -v ON_ERROR_STOP=1 -f migrations/001_init.sql; \
	else \
	  echo "psql not found; using dockerized psql"; \
	  CID=$$(docker compose ps -q db); \
	  docker cp migrations/001_init.sql $$CID:/tmp/001_init.sql; \
	  docker compose exec -T db psql -U $${PG_USER:-mm} -d $${PG_DATABASE:-mmdb} -v ON_ERROR_STOP=1 -f /tmp/001_init.sql; \
	fi

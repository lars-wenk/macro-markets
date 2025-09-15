# Macro Markets – Base Scaffold

### Prereqs

- Go 1.23+
- Docker & Docker Compose
- psql (for migrations)

### 1) Start Infra

```bash
docker compose up -d db nats
```

### 2) Run DB Migration

```bash
make migrate
```

### 3) Start API (local) or via Compose

```bash
# local
HTTP_ADDR=:8080 go run ./cmd/api-gateway
# or docker
docker compose up -d api
```

### 4) Start Demo Ingestor

```bash
docker compose up -d ingestor_x
```

### 5) Try Endpoints

- Health: `GET http://localhost:8080/healthz/live` / `/ready`
- Time: `GET http://localhost:8080/api/v1/time`
- Briefing: `GET http://localhost:8080/api/v1/briefings/latest`
- Calendar: `GET http://localhost:8080/api/v1/calendar/`

> Notes

- JWT is a placeholder; implement real verification before prod.
- Add proper JetStream stream/consumer setup for NATS.
- Replace demo ingestor with your X.com stream/polling worker.

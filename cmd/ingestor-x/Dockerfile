# syntax=docker/dockerfile:1
FROM golang:1.23 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/ingestor-x ./cmd/ingestor-x

 
FROM gcr.io/distroless/base-debian12
WORKDIR /
COPY --from=build /out/ingestor-x /ingestor-x
USER nonroot:nonroot
ENTRYPOINT ["/ingestor-x"]
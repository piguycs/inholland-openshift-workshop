FROM docker.io/library/golang:1.27.1-alpine3.24 AS builder

WORKDIR /app-src

ARG SQLC_VERSION=v1.31.1
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@${SQLC_VERSION}

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN sqlc generate

RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /api ./cmd/api

FROM scratch

COPY --from=builder /api /api

EXPOSE 8080

ENTRYPOINT ["/api"]

package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5"
	"inholland.nl/go-container-template/internal/config"
)

const DefaultAddr string = ":8080"

func main() {
	cfg := config.NewEnvConfigProvider()
	addr := cfg.GetConfig("ADDR", DefaultAddr)

	ctx := context.Background()

	connString := cfg.GetConfig("DATABASE_URL", "")
	if connString == "" {
		panic("DATABASE_URL not set")
	}
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close(ctx)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!\n"))
	})

	mux.HandleFunc("GET /check-sql", func(w http.ResponseWriter, r *http.Request) {
		cmd := conn.QueryRow(r.Context(), "select 1 + 1")

		var out int
		err := cmd.Scan(&out)
		if err != nil {
			w.Write([]byte("Internal server error"))
			slog.Error("Postgres SQL error", "error", err.Error())
			return
		}

		w.Write(fmt.Appendf(nil, "Result of SQL query: 1+1=%d", out))
	})

	slog.Info("Starting HTTP server", "addr", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		slog.Error("HTTP server error", "error", err.Error())
	}
}

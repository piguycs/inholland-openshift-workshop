package main

import (
	"log/slog"
	"net/http"

	"inholland.nl/go-container-template/internal/config"
)

const DefaultAddr string = ":8080"

func main() {
	cfg := config.NewEnvConfigProvider()
	addr := cfg.GetConfig("ADDR", DefaultAddr)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!\n"))
	})

	slog.Info("Starting HTTP server", "addr", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		slog.Error("HTTP server error", "error", err.Error())
	}
}

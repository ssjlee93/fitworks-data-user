package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	pgxinfra "github.com/ssjlee93/fitworks-data-user/internal/pgx"

	userservice "github.com/ssjlee93/fitworks-data-user/core/service/user"

	pgxrepo "github.com/ssjlee93/fitworks-data-user/pkg/postgres"

	httphandler "github.com/ssjlee93/fitworks-data-user/pkg/http"
)

func main() {
	// ── Structured logger ──────────────────────────────────────────────────
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})))

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// ── Database ───────────────────────────────────────────────────────────
	pool, err := pgxinfra.NewPool(ctx)
	if err != nil {
		slog.Error("database init failed", "err", err)
		os.Exit(1)
	}
	defer pool.Close()

	// ── Wire the hexagon ───────────────────────────────────────────────────
	repo := pgxrepo.New(pool)
	svc := userservice.New(repo)
	handler := httphandler.NewHandler(svc)
	router := httphandler.NewRouter(handler)

	// ── HTTP server ────────────────────────────────────────────────────────
	addr := envOrDefault("SERVER_ADDR", ":8080")
	srv := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start listening in a goroutine so we can handle shutdown signals.
	go func() {
		slog.Info("server listening", "addr", addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server error", "err", err)
			os.Exit(1)
		}
	}()

	// Block until signal received, then gracefully shut down.
	<-ctx.Done()
	slog.Info("shutting down…")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("shutdown error", "err", err)
	}
	slog.Info("server stopped")
}

func envOrDefault(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}

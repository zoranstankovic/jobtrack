// Command jobtrack starts the web application.
//
// All settings come from environment variables; see internal/config.
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

	"github.com/labstack/echo/v5"

	"github.com/zoranstankovic/jobtrack/internal/app"
	"github.com/zoranstankovic/jobtrack/internal/config"
)

func main() {
	if err := run(); err != nil {
		slog.Error("jobtrack stopped", "error", err)
		os.Exit(1)
	}
}

func run() error {
	// One structured logger for the whole app (including Echo's request log).
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	// ctx is cancelled when the process receives Ctrl+C (SIGINT) or SIGTERM
	// (what `docker compose down` sends). That triggers a graceful shutdown.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg := config.FromEnv()

	a, err := app.New(ctx, cfg)
	if err != nil {
		return err
	}
	defer a.Close()

	slog.Info("starting jobtrack", "addr", cfg.ListenAddr)
	server := echo.StartConfig{
		Address:         cfg.ListenAddr,
		HideBanner:      true,
		GracefulTimeout: 10 * time.Second,
	}
	// Start blocks until ctx is cancelled, then waits for running requests.
	if err := server.Start(ctx, a.Handler()); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Package app wires jobtrack together: database, migrations and HTTP routes.
//
// Both main and the tests build the app through New, so the tests exercise
// exactly the same startup path as production.
package app

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	"github.com/zoranstankovic/jobtrack/internal/config"
	"github.com/zoranstankovic/jobtrack/internal/db"
	"github.com/zoranstankovic/jobtrack/internal/db/store"
	"github.com/zoranstankovic/jobtrack/internal/web"
)

// App is a fully started jobtrack instance, ready to serve HTTP.
type App struct {
	pool    *pgxpool.Pool
	handler http.Handler
}

// New connects to the database, runs pending migrations and builds the HTTP
// handler. Call Close when done.
func New(ctx context.Context, cfg config.Config) (*App, error) {
	pool, err := db.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Migrate(ctx, pool); err != nil {
		pool.Close()
		return nil, err
	}

	e := echo.New()
	e.Logger = slog.Default()   // same log format as the rest of the app
	e.Use(middleware.Recover()) // turn a panic in a handler into a 500 instead of a crash
	e.Use(middleware.RequestLogger())

	// cfg.BABaseURL is not used yet. Ticket 09 hands it to the BA Source client.
	web.Register(e, store.New(pool))

	return &App{pool: pool, handler: e}, nil
}

// Handler returns the HTTP handler that serves the whole app.
func (a *App) Handler() http.Handler {
	return a.handler
}

// Close releases the database connections.
func (a *App) Close() {
	a.pool.Close()
}

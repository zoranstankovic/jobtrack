// Package config reads jobtrack's settings from environment variables.
//
// Reading configuration from the environment (instead of files or flags) is
// what lets the same Docker image run anywhere: compose, a server, or tests
// each just set different variables.
package config

import "os"

// Default values used when an environment variable is not set. They match a
// local setup where Postgres from docker compose is reachable on localhost.
const (
	DefaultDatabaseURL = "postgres://jobtrack:jobtrack@localhost:5432/jobtrack?sslmode=disable"
	DefaultListenAddr  = ":8080"
	DefaultBABaseURL   = "https://rest.arbeitsagentur.de/jobboerse/jobsuche-service"
)

// Config holds everything the app needs to know about its environment.
type Config struct {
	// DatabaseURL is the Postgres connection string (env: DATABASE_URL).
	DatabaseURL string
	// ListenAddr is the address the HTTP server listens on (env: LISTEN_ADDR).
	ListenAddr string
	// BABaseURL is the base URL of the Bundesagentur für Arbeit API, the only
	// Source (env: BA_BASE_URL). Tests point it at a fake server.
	BABaseURL string
}

// FromEnv builds a Config from environment variables, falling back to the
// defaults above.
func FromEnv() Config {
	return Config{
		DatabaseURL: getenv("DATABASE_URL", DefaultDatabaseURL),
		ListenAddr:  getenv("LISTEN_ADDR", DefaultListenAddr),
		BABaseURL:   getenv("BA_BASE_URL", DefaultBABaseURL),
	}
}

func getenv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}

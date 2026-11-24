package app_test

// This file is the test harness. Every test in this package talks to jobtrack
// the same way a browser (or HTMX) does: over HTTP. Tests never call internal
// functions and never look at SQL; they send requests and check the HTML.
//
// What the harness sets up:
//   - TestMain starts ONE Postgres container for the whole test run.
//   - newTestApp gives each test its OWN empty database inside that container,
//     starts the real app against it (migrations run at startup, as in prod),
//     and serves it with httptest.
//   - Each test also gets a fake BA server (the Source). The app is pointed at
//     it through configuration, exactly like BA_BASE_URL in production.

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"

	"github.com/zoranstankovic/jobtrack/internal/app"
	"github.com/zoranstankovic/jobtrack/internal/config"
)

// adminURL points at the container's default "postgres" database. The harness
// uses it only to create and drop one database per test.
var adminURL string

// dbCounter makes database names unique within one test run.
var dbCounter atomic.Int64

func TestMain(m *testing.M) {
	os.Exit(runTests(m))
}

// runTests is separate from TestMain so that deferred cleanup runs before
// os.Exit (os.Exit skips defers).
func runTests(m *testing.M) int {
	ctx := context.Background()

	pg, err := postgres.Run(ctx, "postgres:18-alpine",
		postgres.WithDatabase("postgres"),
		postgres.WithUsername("jobtrack"),
		postgres.WithPassword("jobtrack"),
		postgres.BasicWaitStrategies(),
	)
	if err != nil {
		log.Printf("starting Postgres container (is Docker running?): %v", err)
		return 1
	}
	defer func() {
		if err := testcontainers.TerminateContainer(pg); err != nil {
			log.Printf("terminating Postgres container: %v", err)
		}
	}()

	adminURL, err = pg.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Printf("reading Postgres connection string: %v", err)
		return 1
	}

	return m.Run()
}

// testApp is one running jobtrack instance with its own database and its own
// fake BA server.
type testApp struct {
	// URL is the base URL of the running app, e.g. "http://127.0.0.1:54321".
	URL string
	// BA is the fake Source. Register handlers on it to control what the app
	// receives when it calls the BA API.
	BA *fakeBA
}

// newTestApp starts a fresh app with an empty, fully migrated database.
// Everything is cleaned up automatically when the test ends.
func newTestApp(t *testing.T) *testApp {
	t.Helper()
	ctx := context.Background()

	dbURL := createTestDatabase(t)
	ba := newFakeBA(t)

	cfg := config.Config{
		DatabaseURL: dbURL,
		ListenAddr:  "unused-in-tests", // httptest picks a free port for us
		BABaseURL:   ba.URL,
	}

	a, err := app.New(ctx, cfg)
	if err != nil {
		t.Fatalf("starting app: %v", err)
	}
	t.Cleanup(a.Close)

	srv := httptest.NewServer(a.Handler())
	t.Cleanup(srv.Close)

	return &testApp{URL: srv.URL, BA: ba}
}

// createTestDatabase creates an empty database for one test and drops it when
// the test ends. Returns its connection URL.
func createTestDatabase(t *testing.T) string {
	t.Helper()
	ctx := context.Background()

	name := fmt.Sprintf("test_%d", dbCounter.Add(1))

	admin, err := pgx.Connect(ctx, adminURL)
	if err != nil {
		t.Fatalf("connecting to Postgres: %v", err)
	}
	defer admin.Close(ctx)

	// CREATE DATABASE does not accept query parameters, so the name is put into
	// the SQL directly. That's safe here because we generate the name ourselves.
	if _, err := admin.Exec(ctx, "CREATE DATABASE "+name); err != nil {
		t.Fatalf("creating test database: %v", err)
	}
	t.Cleanup(func() {
		admin, err := pgx.Connect(ctx, adminURL)
		if err != nil {
			t.Errorf("connecting to Postgres for cleanup: %v", err)
			return
		}
		defer admin.Close(ctx)
		// FORCE closes any connections the app still has open.
		if _, err := admin.Exec(ctx, "DROP DATABASE "+name+" WITH (FORCE)"); err != nil {
			t.Errorf("dropping test database: %v", err)
		}
	})

	u, err := url.Parse(adminURL)
	if err != nil {
		t.Fatalf("parsing admin URL: %v", err)
	}
	u.Path = "/" + name
	return u.String()
}

// get sends a GET request the way a browser would and returns the status code
// and the response body.
func (a *testApp) get(t *testing.T, path string) (int, string) {
	t.Helper()
	res, err := http.Get(a.URL + path)
	if err != nil {
		t.Fatalf("GET %s: %v", path, err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("reading body of GET %s: %v", path, err)
	}
	return res.StatusCode, string(body)
}

// assertContains fails the test if body does not contain every one of wants.
func assertContains(t *testing.T, body string, wants ...string) {
	t.Helper()
	for _, want := range wants {
		if !strings.Contains(body, want) {
			t.Errorf("response does not contain %q\n--- body ---\n%s", want, body)
		}
	}
}

// fakeBA stands in for the Bundesagentur für Arbeit API (the Source).
// Tests register handlers on it, for example:
//
//	app.BA.Handle("GET /pc/v6/jobs", func(w http.ResponseWriter, r *http.Request) { ... })
//
// Any request without a registered handler gets 404, so an unexpected call to
// the Source shows up as an error instead of silently succeeding.
type fakeBA struct {
	URL string
	mux *http.ServeMux
}

func newFakeBA(t *testing.T) *fakeBA {
	t.Helper()
	mux := http.NewServeMux()
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return &fakeBA{URL: srv.URL, mux: mux}
}

// Handle registers a handler for a pattern such as "GET /pc/v6/jobs".
func (f *fakeBA) Handle(pattern string, handler http.HandlerFunc) {
	f.mux.HandleFunc(pattern, handler)
}

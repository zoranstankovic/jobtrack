package app_test

import (
	"net/http"
	"testing"
)

func TestDashboardShowsNavigation(t *testing.T) {
	app := newTestApp(t)

	status, body := app.get(t, "/")

	if status != http.StatusOK {
		t.Fatalf("GET / status = %d, want %d", status, http.StatusOK)
	}
	assertContains(t, body,
		"<title>Dashboard · jobtrack</title>",
		`href="/"`,
		`href="/applications"`,
		`href="/companies"`,
		`href="/search-profiles"`,
		"Search Profiles",
	)
}

func TestPagesLoadHTMX(t *testing.T) {
	app := newTestApp(t)

	_, body := app.get(t, "/")
	assertContains(t, body, `<script src="/static/htmx.min.js"`)

	status, script := app.get(t, "/static/htmx.min.js")
	if status != http.StatusOK {
		t.Fatalf("GET /static/htmx.min.js status = %d, want %d", status, http.StatusOK)
	}
	assertContains(t, script, "htmx")
}

func TestHealthCheckReachesDatabase(t *testing.T) {
	app := newTestApp(t)

	status, body := app.get(t, "/healthz")

	if status != http.StatusOK {
		t.Fatalf("GET /healthz status = %d, want %d (body: %s)", status, http.StatusOK, body)
	}
	assertContains(t, body, "ok")
}

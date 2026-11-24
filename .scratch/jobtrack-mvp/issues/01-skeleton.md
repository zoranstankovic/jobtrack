# 01: Skeleton: app + Postgres in compose, test harness

**What to build:** The user runs `docker compose up` and opens jobtrack in the browser. A placeholder Dashboard page loads with a navigation bar (Dashboard, Applications, Companies, Search Profiles). Postgres runs next to the app, migrations run automatically at startup, and data is kept in a named volume. The test harness is in place: tests start the real Echo app against a real Postgres, send it HTTP requests and check the returned HTML. A fake BA server (`httptest`) can stand in for the Source through configuration. This ticket is written by Claude (boilerplate). The user reads it with an explanation of each part.

**Teaches:** Go modules, Echo routing, `templ` components and layouts, HTMX setup, multi-stage Docker builds, compose with healthchecks and volumes, `goose` migrations, `sqlc` configuration, configuration through environment variables, an integration test harness.

**Blocked by:** None (can start immediately)

**Status:** ready-for-agent

- [ ] `docker compose up` on a clean machine builds the image and starts the app and Postgres; the app waits until the database is healthy
- [ ] The placeholder Dashboard page loads with a shared layout and navigation; HTMX is loaded
- [ ] Migrations run automatically at app startup (an initial empty migration is enough)
- [ ] Data persists across `docker compose down` / `up` (named volume)
- [ ] The database URL, listen address and BA base URL are read from environment variables, with sensible defaults in compose
- [ ] `sqlc` is configured and generates code from SQL queries; `templ` generation is part of the build
- [ ] The test harness starts the app against a real Postgres with a clean state per test, and can point the BA base URL at an `httptest` fake server
- [ ] One passing test requests the Dashboard over HTTP and checks the page content
- [ ] The README explains how to run the app, how to run the tests, and how to regenerate `templ` and `sqlc` code

# jobtrack

A personal tool for tracking job Applications and collecting Job Postings from the Bundesagentur für Arbeit. Built with Go, Echo, templ, HTMX and Postgres. The domain vocabulary is in [`CONTEXT.md`](CONTEXT.md).

## Run it

You need Docker. Then:

```sh
docker compose up --build
```

Open <http://localhost:8080>. Compose starts Postgres first, waits until its healthcheck passes, and then starts the app. The app runs database migrations on startup.

- `docker compose down` stops everything. Your data stays in the `pgdata` volume.
- `docker compose down -v` also deletes the volume, which means **all data**.

### Configuration

All settings come from environment variables. `compose.yaml` sets them for Docker. Running outside Docker uses the defaults.

| Variable       | Default                                                                | Meaning                                    |
| -------------- | ---------------------------------------------------------------------- | ------------------------------------------ |
| `DATABASE_URL` | `postgres://jobtrack:jobtrack@localhost:5432/jobtrack?sslmode=disable` | Postgres connection string                 |
| `LISTEN_ADDR`  | `:8080`                                                                | Address the HTTP server listens on         |
| `BA_BASE_URL`  | `https://rest.arbeitsagentur.de/jobboerse/jobsuche-service`            | Base URL of the BA API (the only Source)   |

### Run the app without Docker (faster while developing)

Start only the database in Docker, then run the app with Go:

```sh
docker compose up -d db
make run          # = go tool templ generate && go run ./cmd/jobtrack
```

## Tests

```sh
make test         # = go tool templ generate && go test ./...
```

Docker must be running. The tests start their own Postgres container through [testcontainers](https://golang.testcontainers.org/), so they don't touch your compose database.

How the tests work (see `internal/app/harness_test.go`):

- Each test calls `newTestApp(t)`. That gives it its own empty database, starts the real app against it (migrations included), and serves the app on a random port.
- Tests talk to the app only over HTTP, the way a browser or HTMX would, and check the returned HTML. They never call internal functions or check SQL.
- Each test also gets a fake BA server (`app.BA`). Register handlers on it to control what the Source returns:

  ```go
  app.BA.Handle("GET /pc/v6/jobs", func(w http.ResponseWriter, r *http.Request) {
      w.Write([]byte(`{"ergebnisliste": []}`))
  })
  ```

## Generated code

Two tools generate Go code. Never edit their output by hand.

| Tool  | Input                                         | Output                        | Regenerate      | In git? |
| ----- | --------------------------------------------- | ----------------------------- | --------------- | ------- |
| templ | `internal/web/views/*.templ`                  | `*_templ.go` next to each file | `make generate` | No      |
| sqlc  | `internal/db/queries/*.sql` + migrations      | `internal/db/store/`          | `make sqlc`     | Yes     |

- **templ** is installed as a Go tool (see the `tool` line in `go.mod`), so `go tool templ` always uses the same version as the library. The Docker build runs it too, which is why its output is not committed. If your editor shows errors about missing `views.Something` functions, run `make generate`.
- **sqlc** runs in Docker (`sqlc/sqlc` image), so you don't need to install it. Its output is committed because the Docker build doesn't run sqlc. Run `make sqlc` after changing a query or a migration.

## Database migrations

Migrations live in `internal/db/migrations/` and are embedded into the binary. [goose](https://github.com/pressly/goose) applies the pending ones every time the app starts.

To change the schema, add a new file with the next number, for example `00002_companies.sql`:

```sql
-- +goose Up
CREATE TABLE companies (...);

-- +goose Down
DROP TABLE companies;
```

Never edit a migration that has already run. Add a new one. Run `make sqlc` afterwards so sqlc knows about the new tables.

## Project layout

```
cmd/jobtrack/          main: reads config, starts the server, graceful shutdown
internal/config/       environment variables → Config
internal/app/          wires everything together (used by main AND tests)
  harness_test.go      the test harness
internal/db/           connection pool + migrations
  migrations/          goose SQL migrations
  queries/             SQL for sqlc
  store/               sqlc output (generated)
internal/web/          Echo routes and handlers
  views/               templ components (layout, pages, fragments)
  static/              htmx.min.js, app.css (embedded into the binary)
```

## Other commands

```sh
make help         # list all make targets
make psql         # psql shell in the compose database
make logs         # follow app logs
```

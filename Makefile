# Common tasks. Run `make help` to list them.

SQLC_VERSION := 1.31.1

.PHONY: help generate sqlc test run up down logs psql

help: ## Show this help
	@grep -E '^[a-z-]+:.*## ' $(MAKEFILE_LIST) | awk -F':.*## ' '{printf "  %-10s %s\n", $$1, $$2}'

generate: ## Generate Go code from .templ files
	go tool templ generate

sqlc: ## Generate Go code from SQL queries (runs sqlc in Docker)
	docker run --rm -v "$(CURDIR)":/src -w /src sqlc/sqlc:$(SQLC_VERSION) generate

test: generate ## Run all tests (needs Docker running)
	go test ./...

run: generate ## Run the app locally against the compose database
	go run ./cmd/jobtrack

up: ## Build and start app + database in Docker
	docker compose up --build

down: ## Stop containers (data is kept)
	docker compose down

logs: ## Follow app logs
	docker compose logs -f app

psql: ## Open a psql shell in the database
	docker compose exec db psql -U jobtrack -d jobtrack

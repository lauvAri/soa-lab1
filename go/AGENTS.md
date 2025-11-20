Repository Guidelines
=====================

Project Structure & Module Organization
---------------------------------------
The repository is polyglot. Primary Go services live in `go/` (sources under `internal/`, entrypoint `main.go`, Gin routes in `routers/`). SQL assets sit in `database/`. Companion Java and Python reference implementations reside in `java/`, `java_gateway/`, and `python/`; treat them as separate apps that mirror the Go feature set. Keep new Go code inside `go/internal/<layer>/` to preserve the current clean layering (controller → service → dao → model).

Build, Test, and Development Commands
-------------------------------------
- `cd go && go run main.go` boots the Gin HTTP service on port 8082. Ensure `.env` is present before running.
- `cd go && go test ./...` exercises all Go packages; set `GOCACHE=/tmp/go-cache` if the default cache is read-only.
- Java modules build with `mvn -f java/pom.xml clean test`. The gateway uses the same command under `java_gateway/`.

Coding Style & Naming Conventions
---------------------------------
Use gofmt for all Go files (`gofmt -w <file>`). File names and package paths must be snake_case (`material_type_dao.go`). Controllers provide `Register...Routes` helpers so routers stay declarative. Favor dependency injection (e.g., DAOs passed into controllers) and keep JSON tags consistent with the existing camelCase API schema.

Testing Guidelines
------------------
Unit tests rely on the standard Go testing package; place new tests beside the code under test with `_test.go` suffixes and descriptive names (e.g., `TestMaterialServiceCreate`). Prefer fast, isolated tests; mock database interactions via DAO/service seams. Java projects use JUnit (via Maven Surefire); Python relies on whatever framework is declared in `python/requirements.txt`.

Commit & Pull Request Guidelines
--------------------------------
Commit messages in history follow short imperative summaries (“Add materials type controller”). Mirror that style and group logically related changes together. Pull requests should describe the problem, the solution, and any testing performed; link to tracking issues when available. Include screenshots or API examples for user-visible changes.

Security & Configuration Tips
-----------------------------
Environment variables load via `go/.env`; never commit secrets. Database DSNs come from `DB_DSN`. When adding third-party dependencies, document them in the relevant README and ensure checksum updates are included (e.g., `go.sum`). For local Docker/MySQL setups, keep schema changes mirrored in `database/material.sql`.***

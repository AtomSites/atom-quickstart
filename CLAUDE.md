# Atom Quickstart

A minimal Go web app template using Echo + Templ.

## Tech Stack

- **Language**: Go 1.25 with Echo v4 framework
- **Templates**: Templ (`.templ` files in `internal/web/components/` and `internal/web/pages/`)
- **Styling**: Single CSS file at `static/css/styles.css`
- **Database**: None
- **Deployment**: Docker multi-stage build, port 3000

## Project Structure

```
cmd/server/main.go          — entry point
internal/config/             — env helpers (GetEnvOrDefault, GetEnvOrPanic, etc.)
internal/web/server.go      — Echo setup, routes, custom error handler
internal/web/pages/          — page handlers + templ templates
internal/web/components/     — shared layout, header, footer
internal/web/middleware/      — custom middleware (security headers)
internal/web/render/         — templ render helper
static/css/styles.css        — all styles
static/assets/               — favicon, images
```

## Dev Workflow

```bash
make install   # install Go deps + templ + reflex + golangci-lint
make dev       # hot-reload dev server on :3000
make build     # compile to bin/server
make test      # generate templ + run all tests
make lint      # run golangci-lint
```

## Adding a New Page

1. Create `internal/web/pages/example.templ` — use `@components.Layout("Title", "example", "Page description") { ... }`
2. Add handler method in `internal/web/pages/handler.go`
3. Register route in `internal/web/server.go`
4. Add nav link in `internal/web/components/header.templ`

## Configuration

Environment variables are loaded via helpers in `internal/config/config.go`:
- `config.GetEnvOrDefault(key, fallback)` — returns value or fallback
- `config.GetEnvOrPanic(key)` — returns value or panics (use for required vars)
- `config.GetEnvOrDefaultInt(key, fallback)` / `config.GetEnvOrDefaultBool(key, fallback)` — typed variants

Copy `.env.example` to `.env` for local development. Never use raw `os.Getenv` — always use the config helpers.

## Error Handling

The app uses a custom `e.HTTPErrorHandler` in `server.go` that renders styled error pages (404 and 500) instead of Echo's default JSON errors. Errors are logged with `slog.Error`.

## UI Development Rules

When adding CSS classes in templ templates, verify they exist in `static/css/styles.css` — add any that are missing.

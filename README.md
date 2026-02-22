# Atom Quickstart

A minimal Go web app template using [Echo](https://echo.labstack.com/) and [Templ](https://templ.guide/). Clone this repo and start building.

## Prerequisites

- [Go 1.25+](https://go.dev/dl/)
- [make](https://www.gnu.org/software/make/)

## Quick Start

```bash
git clone https://github.com/AtomSites/atom-quickstart.git
cd atom-quickstart
cp .env.example .env   # configure environment variables
make install           # installs Go deps, templ, reflex, golangci-lint
make dev               # starts dev server with hot reload on http://localhost:3000
```

## Project Structure

```
cmd/server/main.go              — entry point, graceful shutdown
internal/config/                — env variable helpers
  config.go                     — GetEnvOrDefault, GetEnvOrPanic, typed variants
internal/web/server.go          — Echo setup, routes, error handler
internal/web/pages/             — page handlers + templ templates
  handler.go                    — handler methods (Home, About, NotFound, etc.)
  home.templ                    — home page template
  about.templ                   — about page template
  errors.templ                  — 404 and 500 error page templates
internal/web/components/        — shared UI components
  layout.templ                  — HTML shell with meta tags
  header.templ                  — navigation header
  footer.templ                  — page footer
internal/web/middleware/        — custom middleware
  security.go                   — security response headers
internal/web/render/            — templ render helper
static/css/styles.css           — all styles
static/assets/                  — favicon, images
```

## Adding a New Page

### 1. Create the template

Create `internal/web/pages/example.templ`:

```go
package pages

import "github.com/AtomSites/atom-quickstart/internal/web/components"

templ examplePage() {
    @components.Layout("Example", "example", "Example page description") {
        <main class="hero">
            <div class="container" style="text-align:center;">
                <h1 class="section-title">Example Page</h1>
                <p class="section-subtitle">Your content here.</p>
            </div>
        </main>
    }
}
```

### 2. Add a handler method

In `internal/web/pages/handler.go`:

```go
func (h *Handler) Example(c echo.Context) error {
    return render.Render(c, http.StatusOK, examplePage())
}
```

### 3. Register the route

In `internal/web/server.go`:

```go
e.GET("/example", h.Example)
```

### 4. Add a nav link

In `internal/web/components/header.templ`, add inside the `<nav>`:

```go
<a href="/example" class={ templ.KV("active", activePage == "example") }>Example</a>
```

## Using Atom Components

[Atom Components](https://github.com/AtomSites/atom-components) is an optional library of reusable UI components (modals, forms, cards, toasts) that work with the Atom design system.

### Install

```bash
go get github.com/AtomSites/atom-components@latest
```

### Setup

1. **Serve the component CSS** — add to `internal/web/server.go` after the existing `e.Static` line:

```go
import componentCSS "github.com/AtomSites/atom-components/css"

e.GET("/static/css/atom-components.css", echo.WrapHandler(
    http.StripPrefix("/static/css/", http.FileServer(http.FS(componentCSS.FS))),
))
e.GET("/static/js/atom-components.js", echo.WrapHandler(
    http.StripPrefix("/static/js/", http.FileServer(http.FS(componentCSS.FS))),
))
```

2. **Link the stylesheet** — add to `internal/web/components/layout.templ` after the existing CSS link:

```html
<link rel="stylesheet" href="/static/css/atom-components.css"/>
<script src="/static/js/atom-components.js" defer></script>
```

### Usage

Import and use in any templ file:

```go
import ac "github.com/AtomSites/atom-components"

@ac.Modal("confirm", "Are you sure?") {
    <p>This action cannot be undone.</p>
}

@ac.ContactForm("/contact", ac.ContactFormData{})
```

See the [atom-components README](https://github.com/AtomSites/atom-components) for the full component catalog.

## Make Targets

| Target         | Description                                    |
|----------------|------------------------------------------------|
| `make install` | Install Go dependencies and development tools  |
| `make dev`     | Start dev server with hot reload on :3000      |
| `make build`   | Compile binary to `bin/server`                 |
| `make test`    | Generate templ files and run all tests         |
| `make lint`    | Run golangci-lint                              |
| `make clean`   | Remove build artifacts                         |

## Configuration

Copy `.env.example` to `.env` and edit as needed:

```bash
cp .env.example .env
```

| Variable | Default | Description |
|---|---|---|
| `PORT` | `3000` | Server port |
| `ENVIRONMENT` | `local` | `local`, `staging`, `production` |
| `BASE_URL` | `http://localhost:3000` | Public URL for emails/links |
| `SMTP2GO_API_KEY` | — | SMTP2GO API key |
| `SMTP2GO_SENDER_EMAIL` | — | Sender email address |
| `SMTP2GO_SENDER_NAME` | — | Sender display name |

Use the helpers in `internal/config/config.go` to read env vars — never use raw `os.Getenv`:

```go
port := config.GetEnvOrDefault("PORT", "3000")        // with fallback
dbURL := config.GetEnvOrPanic("DATABASE_URL")          // required, panics if missing
workers := config.GetEnvOrDefaultInt("WORKERS", 4)     // int with fallback
debug := config.GetEnvOrDefaultBool("DEBUG", false)    // bool with fallback
```

## Docker

```bash
docker build -t atom-quickstart .
docker run -p 3000:3000 atom-quickstart
```

The Docker image includes a health check at `/health`.

## AI-Assisted Development

See [CLAUDE.md](CLAUDE.md) for project conventions used by AI coding assistants.

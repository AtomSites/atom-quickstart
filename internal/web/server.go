package web

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	appMiddleware "github.com/AtomSites/atom-quickstart/internal/web/middleware"
	"github.com/AtomSites/atom-quickstart/internal/web/pages"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Recover())
	e.Use(appMiddleware.SecurityHeaders())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
		Skipper: func(c echo.Context) bool {
			path := c.Request().URL.Path
			return strings.HasSuffix(path, ".webp") ||
				strings.HasSuffix(path, ".jpg") ||
				strings.HasSuffix(path, ".jpeg") ||
				strings.HasSuffix(path, ".png") ||
				strings.HasSuffix(path, ".gif") ||
				strings.HasSuffix(path, ".woff2") ||
				strings.HasSuffix(path, ".woff") ||
				strings.HasSuffix(path, ".ico")
		},
	}))

	e.Static("/static", "static")

	h := pages.NewHandler()

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}

		if code == http.StatusNotFound {
			_ = h.NotFound(c)
		} else {
			slog.Error("request error", "status", code, "error", err, "path", c.Request().URL.Path)
			_ = h.ServerError(c)
		}
	}

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	e.GET("/", h.Home)
	e.GET("/about", h.About)

	return &Server{echo: e}
}

func (s *Server) Start(addr string) error {
	return s.echo.Start(addr)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

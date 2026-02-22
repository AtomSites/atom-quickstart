package pages

import (
	"net/http"

	"github.com/AtomSites/atom-quickstart/internal/web/render"
	"github.com/labstack/echo/v4"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Home(c echo.Context) error {
	return render.Render(c, http.StatusOK, homePage())
}

func (h *Handler) About(c echo.Context) error {
	return render.Render(c, http.StatusOK, aboutPage())
}

func (h *Handler) NotFound(c echo.Context) error {
	return render.Render(c, http.StatusNotFound, notFoundPage())
}

func (h *Handler) ServerError(c echo.Context) error {
	return render.Render(c, http.StatusInternalServerError, serverErrorPage())
}

// Package web holds the HTTP layer: Echo routes, handlers and templ views.
//
// Handlers stay thin: they read the request, call the domain code (Tracking,
// Collecting) and render a templ component. There is no JSON API (ADR 0001);
// every response is HTML, either a full page or an HTMX fragment.
package web

import (
	"embed"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"

	"github.com/zoranstankovic/jobtrack/internal/db/store"
	"github.com/zoranstankovic/jobtrack/internal/web/views"
)

//go:embed static
var staticFS embed.FS

// Handlers groups the dependencies that handlers need. New tickets add fields
// here (for example the Tracking module) and register more routes below.
type Handlers struct {
	queries *store.Queries
}

// Register adds all routes to the Echo instance.
func Register(e *echo.Echo, queries *store.Queries) {
	h := &Handlers{queries: queries}

	e.StaticFS("/static", echo.MustSubFS(staticFS, "static"))

	e.GET("/", h.dashboard)
	e.GET("/healthz", h.healthz)
}

func (h *Handlers) dashboard(c *echo.Context) error {
	return render(c, http.StatusOK, views.Dashboard())
}

// healthz reports whether the app can reach its database.
func (h *Handlers) healthz(c *echo.Context) error {
	if _, err := h.queries.Ping(c.Request().Context()); err != nil {
		return c.String(http.StatusServiceUnavailable, "database unreachable")
	}
	return c.String(http.StatusOK, "ok")
}

// render writes a templ component as an HTML response with the given status.
func render(c *echo.Context, status int, component templ.Component) error {
	w := c.Response()
	w.Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
	w.WriteHeader(status)
	return component.Render(c.Request().Context(), w)
}

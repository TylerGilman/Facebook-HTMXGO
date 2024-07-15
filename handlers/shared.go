package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

// Adapter decorater pattern
func Make(h HTTPHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("HTTP handler error", "err", err, "path", r.URL.Path)
		}
	}
}

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}

type contextKey string

const HtmxRequestKey contextKey = "HX-Request"

func setHtmxContext(r *http.Request) *http.Request {
	htmxRequest := r.Header.Get("HX-Request")
	ctx := context.WithValue(r.Context(), HtmxRequestKey, htmxRequest)
	return r.WithContext(ctx)
}

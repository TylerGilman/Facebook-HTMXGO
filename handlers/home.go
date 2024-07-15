package handlers

import (
	"context"
	"log"
	"net/http"

	"facebookhtmx/views/games"
	"facebookhtmx/views/home"
)

type contextKey string

const HtmxRequestKey contextKey = "HX-Request"

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	htmxRequest := r.Header.Get("HX-Request")
	ctx := context.WithValue(r.Context(), HtmxRequestKey, htmxRequest)
	r = r.WithContext(ctx)

	pageMode := r.URL.Query().Get("page_mode")
	log.Printf("Page Mode: %s", pageMode)

	switch pageMode {
	case "about":
		return nil
	case "games":
		return Render(w, r, games.Games())
	default:
		return Render(w, r, home.Index())
	}
}

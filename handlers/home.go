package handlers

import (
	"context"
	"log"
	"net/http"

	"facebookhtmx/views/blog"
	"facebookhtmx/views/friends"
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
	case "friends":
		return Render(w, r, friends.Friends())
	case "games":
		return Render(w, r, games.Games())
	case "blog":
		return Render(w, r, blog.Blog())
	default:
		return Render(w, r, home.Index())
	}
}

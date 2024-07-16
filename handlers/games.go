package handlers

import (
	"facebookhtmx/views/games"
	"net/http"
)

func HandleGames(w http.ResponseWriter, r *http.Request) error {
	isHtmxRequest := r.Header.Get("HX-Request") == "true"
	if isHtmxRequest {
		return Render(w, r, games.Partial())
	} else {
		return Render(w, r, games.Games())
	}
}

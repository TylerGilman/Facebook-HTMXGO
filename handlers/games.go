package handlers

import (
	"facebookhtmx/views/games"
	"log"
	"net/http"
)

func HandleGames(w http.ResponseWriter, r *http.Request) error {
	r = setHtmxContext(r)
	log.Printf("HX-Request: %s", r.Context().Value(HtmxRequestKey))
	return Render(w, r, games.Games())
}

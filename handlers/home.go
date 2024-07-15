package handlers

import (
	"facebookhtmx/views/home"
	"log"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	r = setHtmxContext(r)
	log.Printf("HX-Request: %s", r.Context().Value(HtmxRequestKey))
	return Render(w, r, home.Index())
}
